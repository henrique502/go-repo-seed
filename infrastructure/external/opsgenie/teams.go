package opsgenie

import (
	"github.com/henrique502/go-repo-seed/domain/team"
)

type OpsgenieListTeamResponse struct {
	OpsgenieListResponse
	Data []team.Team `json:"data"`
}

// GetTeamList list teams
func (c *OpsgenieClient) GetTeamList() (*OpsgenieListTeamResponse, error) {
	response := OpsgenieListTeamResponse{}

	_, err := c.client.R().
		SetResult(&response).
		Get("/v2/teams")

	if err != nil {
		return nil, err
	}

	return &response, nil
}
