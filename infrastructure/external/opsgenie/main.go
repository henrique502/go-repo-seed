package opsgenie

import (
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type OpsgenieClient struct {
	client *resty.Client
	host   string
}

type Opsgenie interface {
	// GetAlertList list alerts
	GetAlertList() (*OpsgenieListAlertResponse, error)
	// GetAlertListByDay list alerts
	GetAlertListByDay(day time.Time) (*OpsgenieListAlertResponse, error)
	// GetAlertListPagination list alerts by pagination
	GetAlertListPagination(page string) (*OpsgenieListAlertResponse, error)
	// GetTeamList list teams
	GetTeamList() (*OpsgenieListTeamResponse, error)
	// GetIntegrationList list integrations list
	GetIntegrationsList() (*OpsgenieListIntegrationResponse, error)
}

func New() OpsgenieClient {
	host := os.Getenv("OPSGENIE_HOST")
	token := os.Getenv("OPSGENIE_TOKEN")
	client := resty.New()

	client.SetHostURL(host)
	client.SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"User-Agent":    "Pagar.me",
		"Authorization": "GenieKey " + token,
	})

	return OpsgenieClient{
		client: client,
		host:   host,
	}
}
