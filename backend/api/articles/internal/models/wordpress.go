package models

import (
	"fmt"
	"github.com/wilburt/wilburx9.dev/backend/api/internal"
	"html"
	"regexp"
	"strconv"
	"strings"
)

// WpPost is a container for WordPress post
type WpPost struct {
	ID       int64   `json:"id"`
	Date     string  `json:"date"`
	Modified string  `json:"modified"`
	Link     string  `json:"link"`
	Title    content `json:"title"`
	Excerpt  content `json:"excerpt"`
	Meta     meta    `json:"meta"`
}
type meta struct {
	Thumbnail string `json:"twitter-card-image"`
}

type content struct {
	Rendered string `json:"rendered"`
}

// WpPosts is a slice of WpPost
type WpPosts []WpPost

// ToResult creates ArticleResult by mapping WpPosts to a slice of Article
func (p WpPosts) ToResult(source string) []internal.DbModel {
	var timeLayout = "2006-01-02T15:04:05"
	var articles = make([]internal.DbModel, len(p))

	for i, e := range p {
		articles[i] = Article{
			ID:        internal.MakeId(source, strconv.FormatInt(e.ID, 10)),
			Title:     html.UnescapeString(e.Title.Rendered),
			Thumbnail: e.Meta.Thumbnail,
			Url:       e.Link,
			PostedOn:  internal.StringToTime(timeLayout, e.Date),
			UpdatedOn: internal.StringToTime(timeLayout, e.Date),
			Excerpt:   fmt.Sprintf("%v..", html.UnescapeString(getWpExcept(e.Excerpt.Rendered))),
			Source:    source,
		}
	}
	return articles
}

// Remove Html tag, leading and trailing spaces from the excerpt
func getWpExcept(s string) string {
	// Remove html tags
	var rt = regexp.MustCompile(`<[^>]*>`)
	var noTags = rt.ReplaceAllString(s, " ")

	// Remove double spaces
	var rs = regexp.MustCompile(`/\\s{2,}`)
	var noSpaces = rs.ReplaceAllString(noTags, "")

	return strings.TrimSpace(noSpaces)
}
