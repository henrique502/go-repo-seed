package opsgenie

import (
	"os"

	"github.com/go-resty/resty/v2"
)

var instanceHttp *resty.Client
var host string

func init() {
	host = os.Getenv("OPSGENIE_HOST")
	token := os.Getenv("OPSGENIE_TOKEN")
	instanceHttp = resty.New()

	instanceHttp.SetHostURL(host)
	instanceHttp.SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"User-Agent":    "Pagar.me",
		"Authorization": "GenieKey " + token,
	})
}
