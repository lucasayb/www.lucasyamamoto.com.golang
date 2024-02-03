package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"static_site_generator/loader"
	"static_site_generator/parser"
	"strconv"
	"strings"
)

func GenerateHome(config parser.Config, posts []parser.Post, pages parser.Pages, output string) string {
	createFolder(output)
	from := 0
	to := pages.PerPage

	posts = sortPosts(config, posts)
	postsCount := len(posts)
	var rendered bytes.Buffer
	var fileName string
	for page := 1; page <= pages.PagesCount; page++ {
		if to > postsCount {
			to = postsCount
		}
		postsPage := posts[from:to]
		fmt.Println(posts[0].Frontmatter.Title)
		previousPage := page - 1
		nextPage := page + 1
		paginationData := loader.PaginationData{
			PreviousPage:     previousPage,
			NextPage:         nextPage,
			Page:             page,
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
		to += pages.PerPage
		from += pages.PerPage
	}
	fmt.Println("Home generated successfully")
	return rendered.String()
}

func GeneratePostMultiple(config parser.Config, posts []parser.Post, output string) {
	for _, post := range posts {
		GeneratePost(config, post, output)
	}
	fmt.Println("Posts generated successfully")
}

func GeneratePost(config parser.Config, post parser.Post, output string) {
	createFolder(output)
	postData := loader.PostData{
		Config: config,
		Post:   post,
	}
	content := loader.Post(postData)
	createFile(output, post.Slug, content.Bytes())
}

func GenerateAssets(output string) {
	createFolder(output)
	sourceDir := "static"
	copyFiles(sourceDir, output, sourceDir, output, "")
}

func GenerateJSON(posts []parser.Post) {
	formattedJSON, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("_site/posts.json", formattedJSON, 0666)
}

func sortPosts(config parser.Config, posts []parser.Post) []parser.Post {
	var pivot parser.Post

	sortDirection := config.SortDirection

	for z := 0; z < len(posts); z++ {
		for i := 0; i < len(posts); i++ {
			if sortDirection == "asc" {
				if i+1 >= len(posts) {
					continue
				}
				if posts[i].Frontmatter.Date > posts[i+1].Frontmatter.Date {
					pivot = posts[i]
					posts[i] = posts[i+1]
					posts[i+1] = pivot
				}
			} else {
				if i-1 <= -1 {
					continue
				}
				if posts[i].Frontmatter.Date > posts[i-1].Frontmatter.Date {
					pivot = posts[i]
					posts[i] = posts[i-1]
					posts[i-1] = pivot
				}
			}
		}
	}

	return posts
}

func copyFiles(sourceDir string, outputDir string, baseSourceDir string, baseOutputDir string, baseDir string) {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		log.Fatal(err)
	}
	var filePath string
	baseDir, err = filepath.Abs(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	absBaseSourceDir, err := filepath.Abs(baseSourceDir)
	if err != nil {
		log.Fatal(err)
	}
	absSourceDir, err := filepath.Abs(sourceDir)
	if err != nil {
		log.Fatal(err)
	}
	if baseDir != "" {
		outputDir = strings.Replace(absSourceDir, absBaseSourceDir, "", 1)
	}
	absOutputDir, err := filepath.Abs(filepath.Join(baseDir, baseOutputDir, outputDir))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("----------------------------")
	fmt.Printf("baseDir: %s", baseDir)
	fmt.Println()
	fmt.Printf("sourceDir: %s", sourceDir)
	fmt.Println()
	fmt.Printf("absSourceDir: %s", absSourceDir)
	fmt.Println()
	fmt.Printf("absOutputDir: %s", absOutputDir)
	fmt.Println()
	fmt.Println("----------------------------")
	fmt.Println()
	createFolder(absOutputDir)
	for _, entry := range entries {
		fileName := entry.Name()
		filePath = filepath.Join(absSourceDir, fileName)
		if !entry.IsDir() {
			file, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile(filepath.Join(absOutputDir, fileName), file, 0666)
		} else {
			copyFiles(filePath, outputDir, baseSourceDir, baseOutputDir, baseDir)
		}
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
