package loader

import (
	"bytes"
	"log"
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

func Homepager(homepageData HomepageData) bytes.Buffer {
	tmpl := load("_layouts/home.html")
	var rendered bytes.Buffer
	err := tmpl.ExecuteTemplate(&rendered, "Home", homepageData)
	if err != nil {
		log.Fatal(err)
	}
	return rendered
}

func Post(postData PostData) bytes.Buffer {
	layout := postData.Post.Frontmatter.Layout
	if layout == "" {
		layout = "post"
	}
	tmpl := load("_layouts/" + postData.Post.Frontmatter.Layout + ".html")
	var rendered bytes.Buffer
	err := tmpl.ExecuteTemplate(&rendered, "Post", postData)
	if err != nil {
		log.Fatal(err)
	}
	return rendered
}

func Layouts() *template.Template {
	pattern := "_layouts/*.html"
	return template.Must(template.ParseGlob(pattern))
}

// func Load() {
// 	paginationBytes := Paginate(PaginationData{PreviousLink: "/", NextLink: "/page/1", PagesLength: 10})
// 	fmt.Println(paginationBytes.String())
// }

// func Includes() {
// 	pattern := "_partials/*.html"
// 	return template.Must(template.ParseGlob(pattern))
// }

// func Load(templateFolder string, templateName string) template.Template {
// 	pattern := ""
// 	templates := template.Must(template.ParseGlob(pattern))
// 	// filePath := strings.Join([]string{ templateFolder, "/", templateName, ".html" })

// 	// file, err := os.ReadFile(filePath)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// templateData, errTemplate := template.New(templateName).Parse(string(file))
// 	// if errTemplate != nil {
// 	// 	log.Fatal(errTemplate)
// 	// }
// 	// return templateData.
// }
