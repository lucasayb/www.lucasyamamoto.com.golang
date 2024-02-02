package generator

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"static_site_generator/loader"
	"static_site_generator/parser"
	"strconv"
	"strings"
)

func GenerateHome(config parser.Config, posts []parser.Post, pages parser.Pages, output string) string {
	createFolder(output)
	from := 0
	to := pages.PerPage

	var rendered bytes.Buffer
	var fileName string
	for page := 1; page <= pages.PagesCount; page++ {
		postsPage := posts[from:to]
		previousPage := page - 1
		nextPage := page + 1
		paginationData := loader.PaginationData{
			PreviousPage:     previousPage,
			NextPage:         nextPage,
			Pages:            pages.PagesCount,
			ShowPreviousPage: previousPage >= 1,
			ShowNextPage:     nextPage <= pages.PagesCount,
		}
		homepageData := loader.HomepageData{
			Posts:      postsPage,
			Config:     config,
			Pagination: paginationData,
		}
		rendered = loader.Homepager(homepageData)
		if page == 1 {
			fileName = "index"
		} else {
			fileName = "index-" + strconv.FormatInt(int64(page), 10)
		}
		createFile(output, fileName, rendered.Bytes())
	}
	fmt.Println("Home generated successfully")
	return rendered.String()
}

func GeneratePostMultiple(posts []parser.Post, output string) {
	for _, post := range posts {
		GeneratePost(post, output)
	}
}

func GeneratePost(post parser.Post, output string) {
	createFolder(output)

	contentWithLayout := injectInLayout(post.HTML, post.Layout)

	createFile(output, post.Slug, contentWithLayout)
}

func GenerateAssets() {
	entries, err := os.ReadDir("assets")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func createFolder(output string) {
	_, err := os.Stat(output)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(output, 0755)
		}
	}
}

func createFile(output string, slug string, content []byte) {
	fileName := strings.Join([]string{output, "/", slug, ".html"}, "")
	os.WriteFile(fileName, content, 0755)
}

func injectInLayout(content string, layout string) []byte {
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
	return []byte(strings.Replace(string(file), "{{ content }}", content, 1))
}
