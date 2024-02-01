package main

import (
	"fmt"
	"static_site_generator/parser"
)

func main() {
	files := parser.Parse("_posts")
	fmt.Print(files)
}
