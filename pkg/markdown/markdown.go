package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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

func (c MarkdownContent) Convert() ([]byte, error) {
	var ren = html.WithUnsafe()
	var o = goldmark.WithRendererOptions(ren)
	var ext = goldmark.WithExtensions(extension.Table)
	var m = goldmark.New(ext, o)

	var err error
	var w = &bytes.Buffer{}
	err = m.Convert(c, w)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}