package sync

import (
	"github.com/henrique502/go-repo-seed/config"
	"github.com/henrique502/go-repo-seed/infrastructure/external/opsgenie"
	"github.com/henrique502/go-repo-seed/infrastructure/logger"
)

// Teams fetch and update teams table
func (s SyncApp) Teams() error {
	logger.Log.Infoln("Start teams sync")
	c := opsgenie.New()
	data, err := c.GetTeamList()
	if err != nil {
		return err
	}

	for _, element := range data.Data {
		logger.Log.Infof("Fetch team %s", element.Name)
		config.DB.TeamUpSert(element)
	}

	logger.Log.Infoln("End teams sync")

	return nil
}
