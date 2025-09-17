# Beatrice

A modern, secure, and feature-rich static blog generator written in Go.

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org)
[![Tests](https://img.shields.io/badge/tests-passing-green.svg)](https://github.com/pprobst/beatrice)

https://beatrice-example.surge.sh/

## Features

- 🚀 **Fast Generation**: Generate static HTML from Markdown in seconds
- 🔒 **Secure**: Input validation, path sanitization, and HTML escaping
- 🎨 **Themes**: Built-in light and dark themes
- 📱 **Responsive**: Mobile-friendly design
- ⚡ **Modern Go**: Updated to Go 1.21 with latest dependencies
- 🧪 **Well Tested**: Comprehensive unit test coverage
- 📝 **Rich Logging**: Structured logging with configurable verbosity
- ⚙️ **CLI Options**: Flexible configuration via command-line flags
- 🛡️ **Validation**: Built-in validation for config and content

## Installation

### Option 1: Go Install
```bash
go install github.com/pprobst/beatrice@latest
```

### Option 2: Build from Source
```bash
git clone https://github.com/pprobst/beatrice.git
cd beatrice
go build
```

## Quick Start

1. **Configure your blog:**
   ```bash
   # Edit config.yml
   title: "My Awesome Blog"
   author: "Your Name"
   theme: "dark"  # or "light"
   ```

2. **Create content:**
   ```bash
   # Create posts in the posts/ directory
   # Create about page in about/about.md
   ```

3. **Generate your blog:**
   ```bash
   ./beatrice
   ```

4. **Serve the static files:**
   ```bash
   # Files are generated in the static/ directory
   # Serve with any static file server
   ```

## Usage

### Command Line Options

```bash
./beatrice [options]

Options:
  -config string
        Path to configuration file (default "config.yml")
  -output string
        Output directory for generated files (default "static")
  -verbose
        Enable verbose logging
  -help
        Show help message
```

### Examples

```bash
# Generate with custom config
./beatrice -config my-config.yml

# Generate to custom output directory
./beatrice -output dist/

# Enable verbose logging
./beatrice -verbose

# Combine options
./beatrice -config blog.yml -output public/ -verbose
```

## Configuration

The `config.yml` file supports the following options:

```yaml
title: "Your Blog Title"      # Required: Blog title
author: "Your Name"           # Required: Author name
theme: "dark"                 # Required: "light" or "dark"
```

## Content Structure

### Posts
- Place Markdown files in the `posts/` directory
- Each post must have YAML frontmatter with `title`, `date`, and optional `tags`
- Posts are sorted by the `number` field in descending order

Example post (`posts/my-post.md`):
```markdown
---
number: 1
title: "My First Post"
date: "2023-01-01"
tags:
  - blog
  - introduction
---

# My First Post

This is the content of my first blog post.
```

### About Page
- Create `about/about.md` with your bio/information
- Uses the same Markdown format as posts

## Project Structure

```
beatrice/
├── main.go                    # CLI entry point with logging and flags
├── config.yml                 # Blog configuration
├── src/
│   ├── config/
│   │   ├── config.go          # Configuration loading and validation
│   │   └── config_test.go     # Unit tests for config
│   └── generator/
│       ├── generator.go       # HTML generation with template validation
│       ├── posts.go           # Markdown processing and validation
│       ├── posts_test.go      # Unit tests for posts
│       └── index.go           # Index page data structures
├── posts/                     # Markdown blog posts
├── about/                     # About page content
├── static/                    # Generated output
│   ├── templates/             # HTML templates
│   ├── css/                   # Stylesheets
│   └── *.html                 # Generated pages
└── go.mod                     # Go module dependencies
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./src/config
go test ./src/generator
```

### Building

```bash
go build
```

### Inspirations
- [YASBE](https://github.com/underr/yasbe/)
- [zupzup's blog generator](https://github.com/zupzup/blog-generator)