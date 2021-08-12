package alerts

import (
	"log"
	"time"

	"github.com/henrique502/go-repo-seed/domain"
	"github.com/henrique502/go-repo-seed/infra"
	"github.com/henrique502/go-repo-seed/infra/opsgenie"
	"github.com/henrique502/go-repo-seed/infra/postgre"
)

func syncAlerts(rows []infra.OpsgenieListAlert) {
	for _, element := range rows {
		responderIDs := []string{}

		for _, responder := range element.Responders {
			responderIDs = append(responderIDs, responder.ID)
		}

		alert := domain.Alert{
			ID:            element.ID,
			Message:       element.Message,
			Priority:      element.Priority,
			Source:        element.Source,
			IntegrationID: element.ID,
			ResponderIDs:  responderIDs,
			CreatedAt:     element.CreatedAt,
			UpdatedAt:     element.UpdatedAt,
		}

		postgre.AlertUpSert(alert)
	}
}

func FetchDay(day time.Time) {

	log.Println("Fetch date: " + day.Format("2006-01-02"))
	aletsIterator(day)

}

func aletsIterator(day time.Time) {
	data := opsgenie.GetAlertList(time.Now())
	syncAlerts(data.Data)

	for len(data.Paging.Next) > 0 {
		data = opsgenie.GetAlertListPagination(data.Paging.Next)
		syncAlerts(data.Data)
	}

}
