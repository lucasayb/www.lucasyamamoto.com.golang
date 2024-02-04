package parser

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	markdownParser "github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

type Frontmatter struct {
	Title        string
	Date         string
	Description  string
	Thumbnail    string
	Category     string
	Color        string
	Layout       string
	RedirectFrom []string `yaml:"redirect_from"`
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

type Config struct {
	Title             string
	Email             string
	Description       string
	Baseurl           string
	Url               string
	PerPage           int    `yaml:"per_page"`
	SortDirection     string `yaml:"sort_direction"`
	TwitterUsername   string `yaml:"twitter_username"`
	GithubUsername    string `yaml:"github_username"`
	LinkedinUsername  string `yaml:"linkedin_username"`
	InstagramUsername string `yaml:"instagram_username"`
	DisqusShortname   string `yaml:"disqus_shortname"`
}

func ParseMultiple(config Config, dirName string) ([]Post, Pages) {
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
		parsedPost := Parse(filePath, fileName)
		posts = append(posts, parsedPost)
	}
	perPage := config.PerPage
	postsCount := len(posts)

	pagesCount := math.Ceil(float64(postsCount) / float64(perPage))
	return posts, Pages{PostsCount: postsCount, PerPage: perPage, PagesCount: int(pagesCount)}
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

func ParseConfig() Config {
	config := Config{}
	file, err := os.ReadFile("_config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func parseMarkdown(title string, rawMarkdown []byte) string {
	extensions := markdownParser.CommonExtensions | markdownParser.AutoHeadingIDs
	p := markdownParser.NewWithExtensions(extensions)
	markdownPayload := make([]byte, 0)
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
