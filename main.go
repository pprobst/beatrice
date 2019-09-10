package main

import (
	"github.com/pprobst/beatrice/src/config"
	"github.com/pprobst/beatrice/src/generator"
	//"go get -u gopkg.in/russross/blackfriday.v2"
)

func main() {
	cfg, _ := config.ReadConfig()
	posts := generator.GetPosts(cfg)

	generator.GenerateIndexHTML(cfg, posts)
	generator.GeneratePostsHTML(cfg, posts)
}
