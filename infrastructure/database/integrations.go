package database

import (
	"context"

	"github.com/henrique502/go-repo-seed/domain/integration"
	"github.com/jackc/pgconn"
)

func (c *DatabaseClient) IntegrationUpSert(integration integration.Integration) (pgconn.CommandTag, error) {
	sql := `
    INSERT INTO public.integrations
      (id, name, type, enabled, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, NOW(), NOW())
    ON CONFLICT (id) DO UPDATE SET
      name = excluded.name,
      type = excluded.type,
      enabled = excluded.enabled,
      updated_at = NOW();
  `

	return c.client.Exec(context.Background(), sql,
		integration.ID, integration.Name, integration.Type, integration.Enabled)
}
