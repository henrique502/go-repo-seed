package alerts

import (
	"time"

	"github.com/henrique502/go-repo-seed/domain"
	"github.com/henrique502/go-repo-seed/infra/opsgenie"
	"github.com/henrique502/go-repo-seed/infra/postgre"
)

// SyncByDay fetch and update alerts table by day
func SyncByDay(day time.Time) {
	data := opsgenie.GetAlertListByDay(day)
	syncAlerts(data.Data)

	for len(data.Paging.Next) > 0 {
		data = opsgenie.GetAlertListPagination(data.Paging.Next)
		syncAlerts(data.Data)
	}
}

// Sync fetch and update alerts table
func Sync() {
	data := opsgenie.GetAlertList()
	syncAlerts(data.Data)

	for len(data.Paging.Next) > 0 {
		data = opsgenie.GetAlertListPagination(data.Paging.Next)
		syncAlerts(data.Data)
	}
}

func syncAlerts(rows []domain.AlertOpsgenie) {
	for _, element := range rows {
		responderIDs := []string{}

		for _, responder := range element.Responders {
			responderIDs = append(responderIDs, responder.ID)
		}

		alert := domain.Alert{
			ID:              element.ID,
			Message:         element.Message,
			Priority:        element.Priority,
			Source:          element.Source,
			ReportAckTime:   element.Report.AckTime,
			ReportCloseTime: element.Report.CloseTime,
			IntegrationID:   element.Integration.ID,
			CreatedAt:       element.CreatedAt,
			UpdatedAt:       element.UpdatedAt,
		}

		postgre.AlertUpSert(alert)
	}
}
