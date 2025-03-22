package utils

import (
	"bytes"
	"sync"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
)

// 使用 sync.Pool 复用 buffer
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

// NewMarkdownRenderer 创建一个新的 Markdown 渲染器
func NewMarkdownRenderer(config MarkdownConfig) goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(config.Extensions...),
		goldmark.WithParserOptions(config.ParserOptions...),
		goldmark.WithRendererOptions(config.RendererOptions...),
	)
}

// defaultMarkdownConfig 返回默认的 Markdown 配置
func defaultMarkdownConfig() MarkdownConfig {
	return MarkdownConfig{
		Extensions: []goldmark.Extender{
			extension.Linkify,        // 自动链接支持
			extension.GFM,            // 启用 GitHub Flavored Markdown
			extension.Table,          // 表格支持
			extension.TaskList,       // 任务列表支持
			extension.Strikethrough,  // 删除线支持
			extension.Footnote,       // 脚注支持
			extension.DefinitionList, // 定义列表支持
			extension.Typographer,    // Typography support
			extension.CJK,            // CJK 支持
		},
		ParserOptions: []parser.Option{
			parser.WithAutoHeadingID(),         // 自动生成标题 ID
			parser.WithBlockParsers(),          // 块解析器
			parser.WithInlineParsers(),         // 内联解析器
			parser.WithParagraphTransformers(), // 段落转换器
			parser.WithASTTransformers(),       // AST 转换器
			parser.WithAttribute(),             // 启用自定义属性，目前只有标题支持属性。
		},
		RendererOptions: []renderer.Option{
			html.WithHardWraps(), // 硬换行
			html.WithXHTML(),     // 生成 XHTML
		},
	}
}

// RenderMarkdown 将 Markdown 渲染为 HTML
func RenderMarkdown(content []byte) (string, error) {
	md := NewMarkdownRenderer(defaultMarkdownConfig())
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	if err := md.Convert(content, buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}
