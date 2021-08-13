package opsgenie

import (
	"fmt"
	"strings"
	"time"
)

type OpsgenieListAlertResponse struct {
	OpsgenieListPaginateResponse
	Data []AlertOpsgenie `json:"data"`
}

// GetAlertList list alerts
func (c *OpsgenieClient) GetAlertList() (*OpsgenieListAlertResponse, error) {
	response := OpsgenieListAlertResponse{}
	params := map[string]string{
		"limit":  "100",
		"offset": "0",
		"order":  "desc",
		"sort":   "createdAt",
	}

	_, err := c.client.R().
		SetQueryParams(params).
		SetResult(&response).
		Get("/v2/alerts")

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAlertListByDay list alerts
func (c *OpsgenieClient) GetAlertListByDay(day time.Time) (*OpsgenieListAlertResponse, error) {
	start := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, day.Location())
	end := time.Date(day.Year(), day.Month(), day.Day(), 23, 59, 59, 0, day.Location())

	response := OpsgenieListAlertResponse{}
	params := map[string]string{
		"limit":  "100",
		"offset": "0",
		"order":  "desc",
		"sort":   "createdAt",
		"query":  fmt.Sprintf("createdAt > %d AND createdAt < %d", start.Unix(), end.Unix()),
	}

	_, err := c.client.R().
		SetQueryParams(params).
		SetResult(&response).
		Get("/v2/alerts")

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAlertListPagination list alerts by pagination
func (c *OpsgenieClient) GetAlertListPagination(page string) (*OpsgenieListAlertResponse, error) {
	response := OpsgenieListAlertResponse{}
	url := strings.Replace(page, c.host, "", 1)

	_, err := c.client.R().
		SetResult(&response).
		Get(url)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
