package opsgenie

import (
	"log"

	"github.com/henrique502/go-repo-seed/domain"
)

// GetIntegrationList list integrations list
func GetIntegrationsList() domain.OpsgenieListIntegrationResponse {
	response := domain.OpsgenieListIntegrationResponse{}

	_, err := instanceHttp.R().
		SetResult(&response).
		Get("/v2/integrations")

	if err != nil {
		log.Panic(err)
	}

	return response
}
