package tgmd

import (
	"log"

	"github.com/ai-shift/tgmd/markdownv2"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func Telegramify(s string) string {
	md := []byte(s)
	p := parser.New()
	doc := p.Parse(md)
	opts := markdownv2.RendererOptions{AbsolutePrefix: ""}
	renderer := markdownv2.NewRenderer(opts)
	escaped := markdown.Render(doc, renderer)
	log.Println("Escaped", string(escaped))
	return strng(escaped)
}
