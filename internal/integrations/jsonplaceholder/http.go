package jsonplaceholder

import (
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
func ListPosts() string {
	resp, err := client.Get(baseURL + "/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	return string(body)
}
