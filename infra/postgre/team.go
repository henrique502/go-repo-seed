package postgre

import (
	"context"
	"log"

	"github.com/henrique502/go-repo-seed/domain"
)

func TeamUpSert(team domain.Team) {
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

	_, err := instance.Exec(context.Background(), sql,
		team.ID, team.Name, team.Description)

	if err != nil {
		log.Fatalln(err)
	}
}
