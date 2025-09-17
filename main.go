// Package main provides the CLI interface for the Beatrice static blog generator.
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/pprobst/beatrice/src/config"
	"github.com/pprobst/beatrice/src/generator"
)

func main() {
	configFile := flag.String("config", "config.yml", "Path to configuration file")
	outputDir := flag.String("output", "static", "Output directory for generated files")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	// Configure logging level
	if *verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	if err := generateHTML(*configFile, *outputDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	slog.Info("Blog generation completed successfully")
}

// generateHTML orchestrates the complete blog generation process.
// It validates templates, loads configuration, processes content,
// and generates all HTML files.
func generateHTML(configPath, outputDir string) error {
	slog.Info("Starting blog generation")

	// Validate templates first
	if err := generator.ValidateTemplates(outputDir); err != nil {
		return fmt.Errorf("template validation failed: %v", err)
	}
	slog.Info("Templates validated")

	cfg, err := config.ReadConfigFromFile(configPath)
	if err != nil {
		return fmt.Errorf("ERROR reading config file: %v", err)
	}
	slog.Info("Configuration loaded", "title", cfg.Title)

	posts, err := generator.GetPosts(cfg)
	if err != nil {
		return fmt.Errorf("ERROR getting posts: %v", err)
	}
	slog.Info("Posts processed", "count", len(posts))

	about, err := generator.GetAbout(cfg)
	if err != nil {
		return fmt.Errorf("ERROR getting about page: %v", err)
	}
	slog.Info("About page processed")

	if err := generator.GenerateIndexHTML(cfg, &posts, about, outputDir); err != nil {
		return fmt.Errorf("ERROR generating index HTML: %v", err)
	}
	slog.Info("Index page generated")

	if err := generator.GeneratePostsHTML(cfg, posts, outputDir); err != nil {
		return fmt.Errorf("ERROR generating posts HTML: %v", err)
	}
	slog.Info("All post pages generated")

	return nil
}
