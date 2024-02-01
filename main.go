package main

import (
	"static_site_generator/generator"
	"static_site_generator/parser"
)

func main() {
	output := "_site"
	posts := parser.ParseMultiple("_posts")
	for _, post := range posts {
		generator.Generate(post, output)
	}
}
