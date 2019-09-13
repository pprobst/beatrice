package main

import (
	"fmt"
	"github.com/pprobst/beatrice/src/config"
	"github.com/pprobst/beatrice/src/generator"
)

func main() {
	cfg, _ := config.ReadConfig()
	posts := generator.GetPosts(cfg)

	generator.GenerateIndexHTML(cfg, &posts)
	generator.GeneratePostsHTML(cfg, posts)

	fmt.Println("Fin.")
}
