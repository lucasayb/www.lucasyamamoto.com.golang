package main

import (
	"static_site_generator/generator"
	"static_site_generator/parser"
)

func main() {
	output := "_site"
	posts, pages := parser.ParseMultiple("_posts")
	generator.GenerateHome(posts, pages, output)
	generator.GeneratePostMultiple(posts, output)
}
