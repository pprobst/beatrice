package generator

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pprobst/beatrice/src/config"
	"github.com/russross/blackfriday/v2"
	yaml "gopkg.in/yaml.v2"
)

type Post struct {
	Number    int
	Title     string
	Date      string
	Tags      []string
	Filename  string
	Body      template.HTML
	Theme     string
	BlogTitle string
}

// Validate checks if the post has required fields
func (p *Post) Validate() error {
	if p.Title == "" {
		return fmt.Errorf("post title is required")
	}
	if len(p.Title) > 200 {
		return fmt.Errorf("post title is too long (max 200 characters)")
	}
	if p.Date == "" {
		return fmt.Errorf("post date is required")
	}
	// Basic date format validation
	if len(p.Date) != 10 || p.Date[4] != '-' || p.Date[7] != '-' {
		return fmt.Errorf("post date must be in YYYY-MM-DD format")
	}
	return nil
}

func GetPosts(cfg *config.Config) ([]*Post, error) {
	var posts []*Post
	markdownFiles := getMarkdownFiles("posts")

	// Validate that we have at least one post
	if len(markdownFiles) == 0 {
		return nil, fmt.Errorf("no posts found in posts directory")
	}

	for _, file := range markdownFiles {
		pst, err := readMarkdownFile(file, "posts", cfg)
		if err != nil {
			return nil, fmt.Errorf("ERROR reading post %s: %v", file, err)
		}
		posts = append(posts, pst)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].Number > posts[j].Number
	})

	return posts, nil
}

// The about file is like a post.
func GetAbout(cfg *config.Config) (*Post, error) {
	aboutFiles := getMarkdownFiles("about")
	if len(aboutFiles) == 0 {
		return nil, fmt.Errorf("no about file found in about directory")
	}
	file, err := readMarkdownFile(aboutFiles[0], "about", cfg)
	if err != nil {
		return nil, fmt.Errorf("ERROR reading about file: %v", err)
	}
	return file, nil
}

func getMarkdownFiles(path string) []string {
	var filenames []string
	filesInfo, err := os.ReadDir(path)
	if err != nil {
		return filenames // Return empty slice on error
	}
	for _, file := range filesInfo {
		if filepath.Ext(file.Name()) == ".md" {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames
}

func readMarkdownFile(filename string, folder string, cfg *config.Config) (*Post, error) {
	// Validate filename to prevent directory traversal
	if strings.Contains(filename, "..") || strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		return nil, fmt.Errorf("invalid filename: %s", filename)
	}

	path := filepath.Join(folder, filename)
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("ERROR reading markdown file: %v", err)
	}

	filename = strings.Split(filename, ".")[0]
	lines := strings.Split(string(data), "\n")
	var idx int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "---" && i > 0 {
			idx = i
			break
		}
	}

	// Extract YAML frontmatter
	yamlContent := strings.Join(lines[1:idx], "\n")

	pst := Post{}
	if err := yaml.Unmarshal([]byte(yamlContent), &pst); err != nil {
		return nil, fmt.Errorf("ERROR parsing markdown file: %v", err)
	}

	body := strings.Join(lines[idx+1:], "\n")

	pst.Filename = filename
	pst.Body = template.HTML(blackfriday.Run([]byte(body)))
	pst.BlogTitle = cfg.Title
	pst.Theme = cfg.Theme

	if err := pst.Validate(); err != nil {
		return nil, fmt.Errorf("invalid post %s: %v", filename, err)
	}

	return &pst, nil
}
