package postgre

import (
	"context"
	"log"

	"github.com/henrique502/go-repo-seed/domain"
)

func AlertUpSert(alert domain.Alert) {
	sql := `
    INSERT INTO public.alerts
      (id, priority, source, message, reportAckTime, reportCloseTime, integration_id, responder_ids, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    ON CONFLICT (id) DO UPDATE SET
      name = excluded.name,
      priority = excluded.priority,
      source = excluded.source,
      message = excluded.message,
      reportCloseTime = excluded.reportCloseTime,
      reportAckTime = excluded.reportAckTime,
      integration_id = excluded.integration_id,
      responder_ids = excluded.responder_ids,
      created_at = excluded.created_at,
      updated_at = excluded.updated_at;
  `

	_, err := instance.Exec(context.Background(), sql,
		alert.ID, alert.Name, alert.Priority, alert.Source, alert.Message,
		alert.IntegrationID, alert.ResponderIDs, alert.CreatedAt, alert.UpdatedAt)

	if err != nil {
		log.Fatalln(err)
	}
}
