package models

import (
	"bytes"
	"cmp"
	"embed"
	"encoding/json"
	"log"
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
	Author      string
	PublishedAt time.Time
	Status      string
	HeaderImage Image
	Images      []Image
}
type Image struct {
	Url   string
	Alt   string
	Title string
}

const StatusPublished = "published"
const StatusPreview = "preview"
const StatusDraft = "draft"

func (post *Post) Slug() string {
	if post == nil {
		return "#"
	}
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

//go:embed json
var jsonFiles embed.FS

func GetJsonPosts() []Post {
	var posts []Post
	pathName := "json"
	files, err := jsonFiles.ReadDir(pathName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	for i, entry := range files {
		var filePosts []Post
		log.Printf("entry %d: %+v", i, entry)
		content, err := jsonFiles.ReadFile(strings.Join([]string{pathName, entry.Name()}, "/"))
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
		err = json.Unmarshal(content, &filePosts)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
		posts = append(posts, filePosts...)
	}

	slices.SortFunc(posts, func(i, j Post) int {
		if n := cmp.Compare(i.Day, j.Day); n != 0 {
			return n
		}
		return strings.Compare(i.Title, j.Title)
	})
	return posts
}

func Filter(posts []Post, status string) []*Post {
	filtered := []*Post{}

	for i := range posts {
		if posts[i].Status == status {
			filtered = append(filtered, &posts[i])
		}
	}
	return filtered
}

func LatestPost(posts []*Post) *Post {
	latest := posts[0]
	for _, post := range posts {
		if n := post.PublishedAt.Compare(latest.PublishedAt); n > 0 {
			latest = post
		}
	}
	return latest
}

func FormattedDate(inputTime time.Time, format string) string {

	r := strings.NewReplacer(
		"January", "januari",
		"February", "februari",
		"March", "maart",
		"April", "april",
		"May", "mei",
		"June", "juni",
		"July", "juli",
		"August", "augustus",
		"September", "september",
		"October", "oktober",
		"November", "november",
		"December", "december")

	return r.Replace(inputTime.Format(format))
}
