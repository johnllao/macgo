package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

type MarkdownContent []byte

func (c MarkdownContent) ToHTML() ([]byte, error) {
	var reader = text.NewReader(c)

	var md = goldmark.New(goldmark.WithRendererOptions(html.WithUnsafe()))
	var parser = md.Parser()
	var node = parser.Parse(reader)

	var writer bytes.Buffer
	var renderer = md.Renderer()

	var err error
	err = renderer.Render(&writer, c, node)
	if err != nil {
		return nil, err
	}
	return writer.Bytes(), nil
}