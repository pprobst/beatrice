package generator

import (
	"fmt"
	"github.com/pprobst/beatrice/src/config"
	"gopkg.in/russross/blackfriday.v2"
	yaml "gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
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

func GetPosts(cfg *config.Config) []*Post {
	var posts []*Post
	markdownFiles := getMarkdownFiles("posts")

	for _, file := range markdownFiles {
		pst, _ := readMarkdownFile(file, "posts", cfg)
		posts = append(posts, pst)
	}

	sort.SliceStable(posts, func(i, j int) bool {
		return posts[i].Number > posts[j].Number
	})

	return posts
}

// The about file is like a post.
func GetAbout(cfg *config.Config) *Post {
	aboutFile := getMarkdownFiles("about")[0]
	file, _ := readMarkdownFile(aboutFile, "about", cfg)
	return file
}

func getMarkdownFiles(path string) []string {
	var filenames []string
	filesInfo, _ := ioutil.ReadDir(path)
	for _, file := range filesInfo {
		if filepath.Ext(file.Name()) == ".md" {
			filenames = append(filenames, file.Name())
		}
	}

	return filenames
}

func readMarkdownFile(filename string, folder string, cfg *config.Config) (*Post, error) {
	path := filepath.Join(folder, filename)
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("ERROR reading markdown file: %v", err)
	}

	pst := Post{}
	if err := yaml.Unmarshal(data, &pst); err != nil {
		return nil, fmt.Errorf("ERROR parsing markdown file: %v", err)
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
	body := strings.Join(lines[idx+1:len(lines)], "\n")

	pst.Filename = filename
	pst.Body = template.HTML(blackfriday.Run([]byte(body)))
	pst.BlogTitle = cfg.Title
	pst.Theme = cfg.Theme

	return &pst, nil
}
