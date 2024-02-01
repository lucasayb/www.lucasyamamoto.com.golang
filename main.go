package main

import (
	"static_site_generator/generator"
	"static_site_generator/parser"
)

func main() {
	output := "_site"
	files := parser.ParseMultiple("_posts")
	for _, file := range files {
		generator.Generate(file.Markdown, file.Slug, output)
	}
}
