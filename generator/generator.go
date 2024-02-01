package generator

import (
	"fmt"
	"os"
	"strings"
)

func Generate(content string, slug string, output string) {
	_, err := os.Stat(output)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(output, 0755)
		}
	}
	fileName := strings.Join([]string{output, "/", slug, ".html"}, "")
	fmt.Println(fileName)
	os.WriteFile(fileName, []byte(content), 0755)
	fmt.Println(slug)
}
