package opsgenie

import (
	"os"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client

// Init http opsgenie client
func init() {
	host := os.Getenv("OPSGENIE_HOST")
	token := os.Getenv("OPSGENIE_TOKEN")
	client = resty.New()

	client.SetHostURL(host)
	client.SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"User-Agent":    "Pagar.me",
		"Authorization": "GenieKey " + token,
	})
}
