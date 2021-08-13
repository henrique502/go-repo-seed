package opsgenie

import (
	"github.com/henrique502/go-repo-seed/domain/integration"
)

type OpsgenieListIntegrationResponse struct {
	OpsgenieListResponse
	Data []integration.Integration `json:"data"`
}

// GetIntegrationList list integrations list
func (c *OpsgenieClient) GetIntegrationsList() (*OpsgenieListIntegrationResponse, error) {
	response := OpsgenieListIntegrationResponse{}

	_, err := c.client.R().
		SetResult(&response).
		Get("/v2/integrations")

	if err != nil {
		return nil, err
	}

	return &response, nil
}
