package models

import (
	"bytes"
	"cmp"
	"encoding/json"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/yuin/goldmark"
)

type Post struct {
	Date        time.Time
	Day         int
	Title       string
	Content     string
	HeaderImage Image
	Images      []Image
}
type Image struct {
	Url   string
	Alt   string
	Title string
}

func (post Post) Slug() string {
	return path.Join("dag", strconv.Itoa(post.Day), slug.Make(post.Title))
}

func (post Post) HTML() string {
	// Convert the markdown to HTML, and pass it to the template.
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(post.Content), &buf); err != nil {
		log.Fatalf("failed to convert markdown to HTML: %v", err)
	}

	return buf.String()
}

func GetJsonPosts() []Post {
	var posts []Post
	content, err := os.ReadFile("./posts.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &posts)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	slices.SortFunc(posts, func(i, j Post) int {
		if n := cmp.Compare(i.Day, j.Day); n != 0 {
			return n
		}
		return strings.Compare(i.Title, j.Title)
	})
	return posts
}
