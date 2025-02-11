package render_middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
)

// MarkdownRender 返回 Markdown 渲染中间件
func MarkdownRender() echo.MiddlewareFunc {
	return MarkdownRenderWithConfig(defaultMarkdownConfig())
}

// bodyParser 用于解析 JSON 请求体
type bodyParser struct {
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" form:"content_markdown" query:"content_markdown"`
}

// MarkdownConfig 用于配置 Goldmark 渲染器
type MarkdownConfig struct {
	Extensions      []goldmark.Extender
	ParserOptions   []parser.Option
	RendererOptions []renderer.Option
}

func defaultMarkdownConfig() MarkdownConfig {
	return MarkdownConfig{
		Extensions: []goldmark.Extender{
			extension.GFM, // 启用 GitHub Flavored Markdown
		},
		ParserOptions: []parser.Option{
			parser.WithAutoHeadingID(),         // 自动生成标题 ID
			parser.WithBlockParsers(),          // 块解析器
			parser.WithInlineParsers(),         // 内联解析器
			parser.WithParagraphTransformers(), // 段落转换器
			parser.WithASTTransformers(),       // AST 转换器
		},
		RendererOptions: []renderer.Option{
			html.WithHardWraps(), // 硬换行
			html.WithXHTML(),     // 生成 XHTML
		},
	}
}

// MarkdownRenderWithConfig 返回带有自定义配置的 Markdown 渲染中间件
func MarkdownRenderWithConfig(config MarkdownConfig) echo.MiddlewareFunc {
	md := goldmark.New(
		goldmark.WithExtensions(config.Extensions...),
		goldmark.WithParserOptions(config.ParserOptions...),
		goldmark.WithRendererOptions(config.RendererOptions...),
	)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contentType := c.Request().Header.Get("Content-Type")
			if contentType == "" {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Content-Type 缺失"})
			}

			if strings.HasPrefix(contentType, "multipart/form-data") {
				return handleMultipart(c, md, next)
			} else if contentType == "application/json" {
				return handleJSON(c, md, next)
			} else {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Content-Type 错误"})
			}
		}
	}
}

// handleMultipart 处理表单上传文件
func handleMultipart(c echo.Context, md goldmark.Markdown, next echo.HandlerFunc) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无法解析 multipart 表单"})
	}

	files := form.File["content_markdown"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "文件上传失败, 请检查 content_markdown 参数"})
	}

	file, err := files[0].Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无法读取上传的文件"})
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "无法读取文件内容"})
	}

	return renderMarkdown(c, md, buf.Bytes(), next)
}

// handleJSON 处理 JSON 请求
func handleJSON(c echo.Context, md goldmark.Markdown, next echo.HandlerFunc) error {
	// 限制请求体大小为 20MB
	limitedReader := io.LimitReader(c.Request().Body, 20<<20)
	body := new(bytes.Buffer)
	if _, err := body.ReadFrom(limitedReader); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "读取请求体失败"})
	}

	c.Request().Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	var bodyParser bodyParser
	if err := json.Unmarshal(body.Bytes(), &bodyParser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无法解析 JSON 数据"})
	}

	if bodyParser.ContentMarkdown == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "content_markdown 参数缺失"})
	}

	return renderMarkdown(c, md, []byte(bodyParser.ContentMarkdown), next)
}

// renderMarkdown 将 Markdown 渲染为 HTML 并返回结果
func renderMarkdown(c echo.Context, md goldmark.Markdown, markdown []byte, next echo.HandlerFunc) error {
	var buf bytes.Buffer
	if err := md.Convert(markdown, &buf); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Markdown 渲染失败"})
	}

	c.Set("contentHtml", buf.String())

	return next(c)
}
