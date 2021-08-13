package integrations

import (
	"github.com/henrique502/go-repo-seed/infra/opsgenie"
	"github.com/henrique502/go-repo-seed/infra/postgre"
)

// Sync fetch and update integrations table
func Sync() {
	data := opsgenie.GetIntegrationsList()

	for _, element := range data.Data {
		postgre.IntegrationUpSert(element)
	}
}
