package parser

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
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
	Markdown    string
	Frontmatter Frontmatter
	Layout      string
	Slug        string
	Date        string
}

func ParseMultiple(dirName string) []Post {
	dirEntries, err := os.ReadDir(dirName)
	posts := make([]Post, len(dirEntries))
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
	return posts
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

	post := Post{
		Markdown:    parseMarkdown(rawMarkdown),
		Frontmatter: parseFrontmatter(rawFrontmatter),
	}

	if post.Frontmatter.Layout == "post" {
		post.Layout = "default.html"
	} else {
		post.Layout = strings.Join([]string{post.Frontmatter.Layout, "html"}, ".")
	}
	post.Date = fileName[0:10]
	post.Slug = strings.TrimSuffix(fileName[11:], ".md")

	return post
}

func parseMarkdown(rawMarkdown []byte) string {
	html := markdown.ToHTML(rawMarkdown, nil, nil)
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
