package generator

import (
	"fmt"
	"log"
	"os"
	"static_site_generator/parser"
	"strings"
)

func Generate(post parser.Post, output string) {
	_, err := os.Stat(output)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(output, 0755)
		}
	}

	contentWithLayout := injectInLayout(post.HTML, post.Layout)

	fileName := strings.Join([]string{output, "/", post.Slug, ".html"}, "")
	fmt.Println(fileName)
	os.WriteFile(fileName, []byte(contentWithLayout), 0755)
}

func injectInLayout(content string, layout string) string {
	var layoutFile string
	if layout == "post" || layout == "" {
		layoutFile = "default.html"
	} else {
		layoutFile = strings.Join([]string{layout, "html"}, ".")
	}

	file, err := os.ReadFile(strings.Join([]string{"_layouts", "/", layoutFile}, ""))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(file), "{{ content }}", content, 1)
}
