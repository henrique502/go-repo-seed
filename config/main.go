package config

import (
	"os"

	"github.com/henrique502/go-repo-seed/infrastructure/database"
	"github.com/henrique502/go-repo-seed/infrastructure/external/opsgenie"
	_ "github.com/joho/godotenv/autoload"
)

var (
	DB       database.DatabaseClient
	Opsgenie opsgenie.OpsgenieClient
)

func init() {
	DB = database.New(os.Getenv("POSTGRE_URL"))
	Opsgenie = opsgenie.New()

	defer func() {
		DB.Close()
	}()
}
