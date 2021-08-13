package opsgenie

import (
	"log"

	"github.com/henrique502/go-repo-seed/domain"
)

// GetTeamList list teams
func GetTeamList() domain.OpsgenieListTeamResponse {
	response := domain.OpsgenieListTeamResponse{}

	_, err := instanceHttp.R().
		SetResult(&response).
		Get("/v2/teams")

	if err != nil {
		log.Panic(err)
	}

	return response
}
