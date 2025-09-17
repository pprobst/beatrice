package generator

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/pprobst/beatrice/src/config"
)

func getTemplate(path string) (*template.Template, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("template file does not exist: %s", path)
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR parsing template file %s: %v", path, err)
	}
	return tmpl, nil
}

// ValidateTemplates checks if all required template files exist and are valid
func ValidateTemplates(outputDir string) error {
	requiredTemplates := []string{
		filepath.Join(outputDir, "templates", "index_tmpl.html"),
		filepath.Join(outputDir, "templates", "post_tmpl.html"),
	}

	for _, tmplPath := range requiredTemplates {
		if _, err := getTemplate(tmplPath); err != nil {
			return fmt.Errorf("template validation failed for %s: %v", tmplPath, err)
		}
	}
	return nil
}

func GenerateIndexHTML(cfg *config.Config, posts *[]*Post, about *Post, outputDir string) error {
	tmplPath := filepath.Join(outputDir, "templates", "index_tmpl.html")
	tmpl, err := getTemplate(tmplPath)
	if err != nil {
		return fmt.Errorf("ERROR loading template %s: %v", tmplPath, err)
	}

	filePath := filepath.Join(outputDir, "index.html")
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ERROR creating file %s: %v", filePath, err)
	}
	defer f.Close()

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

func GeneratePostsHTML(cfg *config.Config, posts []*Post, outputDir string) error {
	tmplPath := filepath.Join(outputDir, "templates", "post_tmpl.html")
	tmpl, err := getTemplate(tmplPath)
	if err != nil {
		return fmt.Errorf("ERROR loading template %s: %v", tmplPath, err)
	}

	for _, post := range posts {
		filePath := filepath.Join(outputDir, post.Filename+".html")
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
