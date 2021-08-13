package sync

import (
	"time"

	"github.com/henrique502/go-repo-seed/config"
	"github.com/henrique502/go-repo-seed/domain/alert"
	"github.com/henrique502/go-repo-seed/infrastructure/external/opsgenie"
)

// AlertsByDay fetch and update alerts table by day
func (s SyncApp) AlertsByDay(day time.Time) error {
	c := opsgenie.New()
	data, err := c.GetAlertListByDay(day)
	if err != nil {
		return err
	}
	syncAlerts(data.Data)

	for len(data.Paging.Next) > 0 {
		data, err = c.GetAlertListPagination(data.Paging.Next)
		if err != nil {
			return err
		}
		syncAlerts(data.Data)
	}

	return nil
}

// Alerts fetch and update alerts table
func (s SyncApp) Alerts() error {
	c := opsgenie.New()
	data, err := c.GetAlertList()
	if err != nil {
		return err
	}
	syncAlerts(data.Data)

	for len(data.Paging.Next) > 0 {
		data, err = c.GetAlertListPagination(data.Paging.Next)
		if err != nil {
			return err
		}
		syncAlerts(data.Data)
	}

	return nil
}

func syncAlerts(rows []opsgenie.AlertOpsgenie) {
	for _, element := range rows {
		responderIDs := []string{}

		for _, responder := range element.Responders {
			responderIDs = append(responderIDs, responder.ID)
		}

		alert := alert.Alert{
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

		config.DB.AlertUpSert(alert)
	}
}
