package biz

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type BlogMeta struct {
	Title   string   `json:"title"   yaml:"title"`
	Summary string   `json:"summary" yaml:"summary"`
	Tags    []string `json:"tags"    yaml:"tags"`
	Date    string   `json:"date"    yaml:"date"`
}

type BlogPost struct {
	Slug    string   `json:"slug"`
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
	Date    string   `json:"date"`
	Content string   `json:"content,omitempty"`
}

type BlogUsecase struct {
	dir string
}

func NewBlogUsecase(dir string) *BlogUsecase {
	return &BlogUsecase{dir: dir}
}

func (b *BlogUsecase) ListPosts() ([]BlogPost, error) {
	entries, err := os.ReadDir(b.dir)
	if err != nil {
		if os.IsNotExist(err) {
			return []BlogPost{}, nil
		}
		return nil, fmt.Errorf("read blog dir: %w", err)
	}

	var posts []BlogPost
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		slug := strings.TrimSuffix(entry.Name(), ".md")
		raw, err := os.ReadFile(filepath.Join(b.dir, entry.Name()))
		if err != nil {
			continue
		}
		meta, _ := parseFrontmatter(string(raw))
		posts = append(posts, BlogPost{
			Slug:    slug,
			Title:   meta.Title,
			Summary: meta.Summary,
			Tags:    meta.Tags,
			Date:    meta.Date,
		})
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})
	return posts, nil
}

func (b *BlogUsecase) GetPost(slug string) (*BlogPost, error) {
	if !isValidSlug(slug) {
		return nil, ErrValidation
	}
	raw, err := os.ReadFile(filepath.Join(b.dir, slug+".md"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("read post: %w", err)
	}
	meta, body := parseFrontmatter(string(raw))
	return &BlogPost{
		Slug:    slug,
		Title:   meta.Title,
		Summary: meta.Summary,
		Tags:    meta.Tags,
		Date:    meta.Date,
		Content: body,
	}, nil
}

func (b *BlogUsecase) CreatePost(slug, content string) error {
	if !isValidSlug(slug) {
		return ErrValidation
	}
	path := filepath.Join(b.dir, slug+".md")
	if _, err := os.Stat(path); err == nil {
		return ErrAlreadyExists
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func (b *BlogUsecase) UpdatePost(slug, content string) error {
	if !isValidSlug(slug) {
		return ErrValidation
	}
	path := filepath.Join(b.dir, slug+".md")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrNotFound
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func (b *BlogUsecase) DeletePost(slug string) error {
	if !isValidSlug(slug) {
		return ErrValidation
	}
	path := filepath.Join(b.dir, slug+".md")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrNotFound
	}
	return os.Remove(path)
}

func parseFrontmatter(raw string) (BlogMeta, string) {
	var meta BlogMeta
	content := raw

	if !strings.HasPrefix(raw, "---\n") {
		return meta, content
	}
	end := strings.Index(raw[4:], "\n---")
	if end < 0 {
		return meta, content
	}
	fmBlock := raw[4 : 4+end]
	content = strings.TrimSpace(raw[4+end+4:])

	_ = yaml.Unmarshal([]byte(fmBlock), &meta)
	if meta.Date == "" {
		meta.Date = time.Now().Format("2006-01-02")
	}
	if meta.Tags == nil {
		meta.Tags = []string{}
	}
	return meta, content
}

func isValidSlug(slug string) bool {
	if slug == "" || len(slug) > 128 {
		return false
	}
	for _, c := range slug {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '-' || c == '_') {
			return false
		}
	}
	return true
}
