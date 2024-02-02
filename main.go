package main

import (
	"static_site_generator/generator"
	"static_site_generator/parser"
)

func main() {
	output := "_site"
	config := parser.ParseConfig()
	posts, pages := parser.ParseMultiple(config, "_posts")
	generator.GenerateAssets()
	generator.GenerateHome(config, posts, pages, output)
	generator.GeneratePostMultiple(posts, output)
}
