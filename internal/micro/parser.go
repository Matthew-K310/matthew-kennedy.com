package micro

import (
	"io/fs"
	"time"

	"github.com/niklasfasching/go-org/org"
)

type Post struct {
	Title   string
	Date    time.Time
	Slug    string
	Content string
}

func LoadPosts(fsys fs.FS, path string) ([]Post, error) {
	f, err := fsys.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc := org.New().Parse(f, path)
	var posts []Post

	for _, node := range doc.Nodes {
		h, ok := node.(org.Headline)
		if !ok || h.Lvl != 1 {
			continue
		}

		var dateStr, slug string
		if h.Properties != nil {
			dateStr, _ = h.Properties.Get("DATE")
			slug, _ = h.Properties.Get("SLUG")
		}

		title := org.String(h.Title...)
		date, _ := time.Parse("2006-01-02 15:04", dateStr)

		w := org.NewHTMLWriter()
		content := w.WriteNodesAsString(h.Children...)

		posts = append(posts, Post{
			Title:   title,
			Date:    date,
			Slug:    slug,
			Content: content,
		})
	}

	// Sort newest first
	for i, j := 0, len(posts)-1; i < j; i, j = i+1, j-1 {
		posts[i], posts[j] = posts[j], posts[i]
	}

	return posts, nil
}
