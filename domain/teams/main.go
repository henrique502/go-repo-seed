package teams

import (
	"github.com/henrique502/go-repo-seed/infra/opsgenie"
	"github.com/henrique502/go-repo-seed/infra/postgre"
)

// Sync fetch and update teams table
func Sync() {
	data := opsgenie.GetTeamList()

	for _, element := range data.Data {
		postgre.TeamUpSert(element)
	}
}
