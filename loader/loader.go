package loader

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
	"static_site_generator/parser"
	"text/template"
)

type HomepageData struct {
	Posts      []parser.Post
	Config     parser.Config
	Pagination PaginationData
}

type PostData struct {
	Config parser.Config
	Post   parser.Post
}

type PaginationData struct {
	PreviousPage     int
	NextPage         int
	Pages            int
	Page             int
	ShowPreviousPage bool
	ShowNextPage     bool
}

func load(filePath string) *template.Template {
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.ParseGlob("_includes/*.html")
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

func loadAds() []byte {
	file, err := os.OpenFile("_includes/ads.html", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var fileData []byte
	for scanner.Scan() {
		line := scanner.Bytes()
		fileData = append(fileData, line...)
	}

	return fileData
}

func Homepager(homepageData HomepageData) bytes.Buffer {
	tmpl := load("_layouts/home.html")
	var rendered bytes.Buffer
	err := tmpl.ExecuteTemplate(&rendered, "home", homepageData)
	if err != nil {
		log.Fatal(err)
	}
	return rendered
}

func Post(postData PostData) bytes.Buffer {
	layout := postData.Post.Frontmatter.Layout

	postData.Post.HTML = injectAds(postData.Post.HTML)

	if layout == "" {
		layout = "post"
	}
	tmpl := load("_layouts/" + layout + ".html")
	var rendered bytes.Buffer
	err := tmpl.ExecuteTemplate(&rendered, layout, postData)
	if err != nil {
		log.Fatal(err)
	}
	return rendered
}

func injectAds(html string) string {
	ads := loadAds()
	matcher := regexp.MustCompile("<!-- ADS -->")
	return string(matcher.ReplaceAll([]byte(html), ads))
}

func Layouts() *template.Template {
	pattern := "_layouts/*.html"
	return template.Must(template.ParseGlob(pattern))
}
