package parser

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	markdownParser "github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

type Frontmatter struct {
	Title         string
	Date          string
	Description   string
	Thumbnail     string
	Category      string
	Color         string
	Layout        string
	Redirect_from []string
}

type Post struct {
	HTML        string
	Frontmatter Frontmatter
	Layout      string
	Slug        string
	Date        string
}

type Pages struct {
	PostsCount int
	PerPage    int
	PagesCount int
}

func ParseMultiple(dirName string) ([]Post, Pages) {
	dirEntries, err := os.ReadDir(dirName)
	posts := make([]Post, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, dirEntry := range dirEntries {
		fileName := dirEntry.Name()
		if !strings.HasSuffix(fileName, ".md") {
			continue
		}
		filePath := strings.Join([]string{dirName, fileName}, "/")

		posts = append(posts, Parse(filePath, fileName))
	}
	perPage := 10
	postsCount := len(posts)
	pagesCount := postsCount / perPage
	return posts, Pages{PostsCount: postsCount, PerPage: perPage, PagesCount: pagesCount}
}

func Parse(filePath string, fileName string) Post {
	file, errRead := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if errRead != nil {
		log.Fatal(errRead)
	}
	scanner := bufio.NewScanner(file)
	var inFrontmatterContext bool = false
	var frontmatterExtracted bool = false
	var rawFrontmatter []byte
	var rawMarkdown []byte
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" && !frontmatterExtracted {
			if inFrontmatterContext {
				inFrontmatterContext = false
				frontmatterExtracted = true
				/* End of the frontmatter*/
				continue
			}

			/* Beginning of the frontmatter */
			inFrontmatterContext = true
			continue
		}
		if inFrontmatterContext {
			rawFrontmatter = append(rawFrontmatter, line...)
			rawFrontmatter = append(rawFrontmatter, '\n')
		} else {
			rawMarkdown = append(rawMarkdown, line...)
			rawMarkdown = append(rawMarkdown, '\n')
		}
	}

	file.Close()

	if !frontmatterExtracted {
		log.Fatal("Frontmatter malformatted.")
	}

	parsedFrontmatter := parseFrontmatter(rawFrontmatter)

	post := Post{
		Frontmatter: parsedFrontmatter,
		HTML:        parseMarkdown(parsedFrontmatter.Title, rawMarkdown),
	}

	post.Date = fileName[0:10]
	post.Slug = strings.TrimSuffix(fileName[11:], ".md")

	return post
}

func parseMarkdown(title string, rawMarkdown []byte) string {
	extensions := markdownParser.CommonExtensions | markdownParser.AutoHeadingIDs
	p := markdownParser.NewWithExtensions(extensions)
	markdownPayload := make([]byte, 0)
	markdownPayload = append(markdownPayload, []byte(strings.Join([]string{"#", title}, " "))...)
	markdownPayload = append(markdownPayload, '\n')
	markdownPayload = append(markdownPayload, rawMarkdown...)
	html := markdown.ToHTML(markdownPayload, p, nil)
	return string(html)
}

func parseFrontmatter(rawFrontmatter []byte) Frontmatter {
	frontmatter := Frontmatter{}
	err := yaml.Unmarshal(rawFrontmatter, &frontmatter)
	if err != nil {
		log.Fatal(err)
	}
	return frontmatter
}
