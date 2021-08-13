package postgre

import (
	"context"
	"log"

	"github.com/henrique502/go-repo-seed/domain"
)

func AlertUpSert(alert domain.Alert) {
	sql := `
    INSERT INTO public.alerts
      (id, priority, source, message, report_ack_time, report_close_time, integration_id, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    ON CONFLICT (id) DO UPDATE SET
      priority = excluded.priority,
      source = excluded.source,
      message = excluded.message,
      report_close_time = excluded.report_close_time,
      report_ack_time = excluded.report_ack_time,
      integration_id = excluded.integration_id,
      created_at = excluded.created_at,
      updated_at = excluded.updated_at;
  `

	_, err := instance.Exec(context.Background(), sql,
		alert.ID, alert.Priority, alert.Source, alert.Message, alert.ReportAckTime, alert.ReportCloseTime,
		alert.IntegrationID, alert.CreatedAt, alert.UpdatedAt)

	if err != nil {
		log.Fatalln(err)
	}
}
