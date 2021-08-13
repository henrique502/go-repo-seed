package sync

import (
	"github.com/henrique502/go-repo-seed/config"
	"github.com/henrique502/go-repo-seed/infrastructure/external/opsgenie"
)

// Teams fetch and update teams table
func (s SyncApp) Teams() error {
	c := opsgenie.New()
	data, err := c.GetTeamList()
	if err != nil {
		return err
	}

	for _, element := range data.Data {
		config.DB.TeamUpSert(element)
	}

	return nil
}
