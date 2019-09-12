package generator

import (
	"bufio"
	"fmt"
	"github.com/pprobst/beatrice/src/config"
	"html/template"
	"os"
	"path/filepath"
)

func getTemplate(path string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR reading template file %s: %v", path, err)
	}
	return tmpl, nil
}

func GenerateIndexHTML(cfg *config.Config, posts *[]*Post) error {
	tmplPath := filepath.Join("static", "index_tmpl.html")
	tmpl, _ := getTemplate(tmplPath)

	filePath := filepath.Join("static", "index.html")
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ERROR creating file %s: %v", filePath, err)
	}
	defer f.Close()

    about := GetAbout(cfg)
	data := IndexPage{
		BlogTitle: cfg.Title,
		Author:    cfg.Author,
		Posts:     *posts,
        About:     about,
		Theme:     cfg.Theme,
		Descr:     cfg.Description,
	}

    // about is treated like a post; it'll be generated like the others
    *posts = append(*posts, about)

	w := bufio.NewWriter(f)

	if err := tmpl.Execute(w, data); err != nil {
		return fmt.Errorf("ERROR executing template %s: %v", filePath, err)
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("ERROR writing file %s: %v", filePath, err)
	}

	return nil
}

func GeneratePostsHTML(cfg *config.Config, posts []*Post) error {
	tmplPath := filepath.Join("static", "post_tmpl.html")
	tmpl, _ := getTemplate(tmplPath)

	for _, post := range posts {
		filePath := filepath.Join("static", post.Filename+".html")
		f, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("ERROR creating file %s: %v", filePath, err)
		}
		defer f.Close()

		w := bufio.NewWriter(f)

		if err := tmpl.Execute(w, post); err != nil {
			return fmt.Errorf("ERROR executing template %s: %v", filePath, err)
		}
		if err := w.Flush(); err != nil {
			return fmt.Errorf("ERROR writing file %s: %v", filePath, err)
		}
	}
	return nil
}
