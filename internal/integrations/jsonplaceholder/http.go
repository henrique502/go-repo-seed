package jsonplaceholder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	once    sync.Once
	client  *http.Client
	baseURL string
)

type JSONPlaceholderPost struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Init init http client instance
func Init() {
	once.Do(func() {
		baseURL = os.Getenv("JSONPLACEHOLDER_URL")
		client = &http.Client{
			Timeout: time.Second * 20,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   20 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		}
	})
}

// ListPosts list posts
func ListPosts() []JSONPlaceholderPost {
	r, getErr := client.Get(baseURL + "/posts")
	if getErr != nil {
		log.Fatalln(getErr)
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	posts := []JSONPlaceholderPost{}
	jsonErr := json.Unmarshal(body, &posts)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return posts
}
