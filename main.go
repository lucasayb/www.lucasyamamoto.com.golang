package main

import (
	"static_site_generator/generator"
	"static_site_generator/parser"
)

func main() {
	output := "_site"
	config := parser.ParseConfig()
	entries, pagination := parser.ParseMultiple(config, "_posts")
	generator.GenerateAssets()
	generator.GenerateHome(config, entries, pagination, output)
	generator.GeneratePostMultiple(config, entries, output)

	pages, _ := parser.ParseMultiple(config, "_pages")
	generator.GeneratePostMultiple(config, pages, output)
}
