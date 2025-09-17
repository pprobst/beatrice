# Beatrice

A barebones static blog generator written in Go.

https://beatrice-example.surge.sh/

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
│   └── generator/
│       ├── generator.go       # HTML generation with template validation
│       ├── posts.go           # Markdown processing and validation
│       └── index.go           # Index page data structures
├── posts/                     # Markdown blog posts
├── about/                     # About page content
├── static/                    # Generated output
│   ├── templates/             # HTML templates
│   ├── css/                   # Stylesheets
│   └── *.html                 # Generated pages
└── go.mod                     # Go module dependencies
```

### Inspirations
- [YASBE](https://github.com/underr/yasbe/)
- [zupzup's blog generator](https://github.com/zupzup/blog-generator)
