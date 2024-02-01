package parser_test

import (
	"static_site_generator/parser"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	parser.ParseMultiple("_posts")
}
