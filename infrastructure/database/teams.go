package database

import (
	"context"

	"github.com/henrique502/go-repo-seed/domain/team"
	"github.com/jackc/pgconn"
)

func (c *DatabaseClient) TeamUpSert(team team.Team) (pgconn.CommandTag, error) {
	sql := `
    INSERT INTO teams
      (id, name, description, created_at, updated_at)
    VALUES
      ($1, $2, $3, NOW(), NOW())
    ON CONFLICT (id) DO UPDATE SET
      name = excluded.name,
      description = excluded.description,
      updated_at = NOW();
  `

	return c.client.Exec(context.Background(), sql,
		team.ID, team.Name, team.Description)
}
