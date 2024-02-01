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
	Redirect_from []string
}

type ParsedFile struct {
	markdown    string
	frontmatter Frontmatter
}

func Parse(dirName string) []ParsedFile {
	return filesReader(dirName)
}

func filesReader(dirName string) []ParsedFile {
	dirEntries, err := os.ReadDir(dirName)
	parsedFiles := make([]ParsedFile, len(dirEntries))
	if err != nil {
		log.Fatal(err)
	}

	for _, dirEntry := range dirEntries {
		parsedFile := ParsedFile{}
		file_path := strings.Join([]string{dirName, dirEntry.Name()}, "/")
		file, errRead := os.OpenFile(file_path, os.O_RDONLY, 0755)
		if errRead != nil {
			log.Fatal(errRead)
		}
		parsedFile.markdown, parsedFile.frontmatter = fileParser(file)
		parsedFiles = append(parsedFiles, parsedFile)
		file.Close()
	}
	return parsedFiles
}

func fileParser(file *os.File) (string, Frontmatter) {
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

	return parseMarkdown(rawMarkdown), parseFrontmatter(rawFrontmatter)
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
