package main

import (
	"fmt"
	"github.com/pprobst/beatrice/src/config"
	"github.com/pprobst/beatrice/src/generator"
	"os"
)

func main() {
	if err := generateHTML(); err != nil {
		fmt.Errorf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Println("Fin.")
}

func generateHTML() error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("ERROR reading config file: %v", err)
	}
	posts := generator.GetPosts(cfg)

	if err := generator.GenerateIndexHTML(cfg, &posts); err != nil {
		return fmt.Errorf("ERROR generating index HTML: %v", err)
	}
	if err := generator.GeneratePostsHTML(cfg, posts); err != nil {
		return fmt.Errorf("ERROR generating posts HTML: %v", err)
	}

	return nil
}
