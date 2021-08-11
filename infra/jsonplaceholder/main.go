package jsonplaceholder

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/henrique502/go-repo-seed/infra"
)

var client *resty.Client

// Init http jsonplaceholder client
func init() {
	host := os.Getenv("JSONPLACEHOLDER_URL")
	client = resty.New()

	client.SetHostURL(host)
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "Pagar.me",
	})
}

// ListPosts list posts
func ListPosts() []infra.JSONPlaceholderPost {
	posts := []infra.JSONPlaceholderPost{}
	_, err := client.R().
		SetResult(&posts).
		Get("/posts")
	if err != nil {
		log.Panic(err)
	}

	return posts
}
