package renderMiddleware

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// bodyParser 用于解析 JSON 请求体
type bodyParser struct {
	ContentMarkdown string `json:"content_markdown" xml:"content_markdown" query:"content_markdown" validator:"required" default:""`
}

// 使用sync.Pool来复用buffer
var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// MarkdownConfig 用于配置 Goldmark 渲染器
type MarkdownConfig struct {
	Extensions      []goldmark.Extender
	ParserOptions   []parser.Option
	RendererOptions []renderer.Option
}

// MarkdownRender 返回 Markdown 渲染中间件
func MarkdownRender() echo.MiddlewareFunc {
	return MarkdownRenderWithConfig(defaultMarkdownConfig())
}

// defaultMarkdownConfig 返回默认的 Markdown 配置
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

			if contentType == "application/json" {
				return handleJSON(c, md, next)
			}

			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Content-Type 错误"})
		}
	}
}

// handleJSON 处理 JSON 请求
func handleJSON(c echo.Context, md goldmark.Markdown, next echo.HandlerFunc) error {
	body := bufferPool.Get().(*bytes.Buffer)
	body.Reset()
	defer bufferPool.Put(body)

	// 限制请求体大小为 20MB
	limitedReader := io.LimitReader(c.Request().Body, 20<<20)
	if _, err := body.ReadFrom(limitedReader); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "读取请求体失败"})
	}

	c.Request().Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	var bodyParser bodyParser
	if err := json.NewDecoder(bytes.NewReader(body.Bytes())).Decode(&bodyParser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "无法解析 JSON 数据"})
	}

	// 判断内容是否为文件路径
	var content string
	if isFilePath(bodyParser.ContentMarkdown) {
		fileContent, err := readFileContent(bodyParser.ContentMarkdown)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "读取文件失败"})
		}
		content = fileContent
	} else {
		content = bodyParser.ContentMarkdown
	}

	// 使用自定义上下文键以避免命名冲突
	c.Set("_temp_content_markdown", content)
	err := renderMarkdown(c, md, []byte(content), next)

	// 清理临时数据
	c.Set("_temp_content_markdown", nil)

	return err
}

// renderMarkdown 将 Markdown 渲染为 HTML 并返回结果
func renderMarkdown(c echo.Context, md goldmark.Markdown, content []byte, next echo.HandlerFunc) error {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	if err := md.Convert(content, buf); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Markdown 渲染失败"})
	}

	// 使用自定义上下文键存储HTML内容
	c.Set("_temp_content_html", buf.String())

	return next(c)
}

// isFilePath 判断是否为文件路径
func isFilePath(content string) bool {
	osType := runtime.GOOS
	if osType == "windows" {
		return (strings.Contains(content, ":\\") && len(content) > 3 && filepath.Ext(content) != "") ||
			(strings.HasPrefix(content, "\\") && len(content) > 1 && filepath.Ext(content) != "")
	}
	return strings.HasPrefix(content, "/") && len(content) > 1 && filepath.Ext(content) != ""
}

// readFileContent 读取文件内容
func readFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
