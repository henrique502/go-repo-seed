package opsgenie

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/henrique502/go-repo-seed/domain"
)

// GetAlertList list alerts
func GetAlertList() domain.OpsgenieListAlertResponse {
	response := domain.OpsgenieListAlertResponse{}
	params := map[string]string{
		"limit":  "100",
		"offset": "0",
		"order":  "desc",
		"sort":   "createdAt",
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

// GetAlertListByDay list alerts
func GetAlertListByDay(day time.Time) domain.OpsgenieListAlertResponse {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())

	response := domain.OpsgenieListAlertResponse{}
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

// GetAlertListPagination list alerts by pagination
func GetAlertListPagination(page string) domain.OpsgenieListAlertResponse {
	response := domain.OpsgenieListAlertResponse{}
	url := strings.Replace(page, host, "", 1)

	_, err := instanceHttp.R().
		SetResult(&response).
		Get(url)

	if err != nil {
		log.Panic(err)
	}

	return response
}
