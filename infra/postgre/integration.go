package postgre

import (
	"context"
	"log"

	"github.com/henrique502/go-repo-seed/domain"
)

func IntegrationUpSert(integration domain.Integration) {
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

	_, err := instance.Exec(context.Background(), sql,
		integration.ID, integration.Name, integration.Type, integration.Enabled)

	if err != nil {
		log.Fatalln(err)
	}
}
