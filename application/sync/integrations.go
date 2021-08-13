package sync

import (
	"github.com/henrique502/go-repo-seed/config"
	"github.com/henrique502/go-repo-seed/infrastructure/external/opsgenie"
)

// Integrations fetch and update integrations table
func (s SyncApp) Integrations() error {
	c := opsgenie.New()
	data, err := c.GetIntegrationsList()
	if err != nil {
		return err
	}

	for _, element := range data.Data {
		config.DB.IntegrationUpSert(element)
	}

	return nil
}
