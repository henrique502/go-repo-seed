package opsgenie

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/henrique502/go-repo-seed/infra"
)

var instanceHttp *resty.Client
var host string

// Init http opsgenie client
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

func GetAlertList(day time.Time) infra.OpsgenieListAlertResponse {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())

	response := infra.OpsgenieListAlertResponse{}
	params := map[string]string{
		"limit":  "100",
		"offset": "0",
		"order":  "desc",
		"sort":   "createdAt",
		"query":  fmt.Sprintf("createdAt > %d AND createdAt < %d", start.Unix(), end.Unix()),
	}

	_, err := instanceHttp.R().
		SetQueryParams(params).
		SetResult(&response).
		Get("/v2/alerts")

	if err != nil {
		log.Panic(err)
	}

	return response
}

func GetAlertListPagination(page string) infra.OpsgenieListAlertResponse {
	log.Println("aaa")
	response := infra.OpsgenieListAlertResponse{}
	url := strings.Replace(page, host, "", 1)
	log.Println("new url" + url)

	_, err := instanceHttp.R().
		SetResult(&response).
		Get(url)

	if err != nil {
		log.Panic(err)
	}

	return response
}
