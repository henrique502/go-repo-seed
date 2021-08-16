package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/henrique502/go-repo-seed/infrastructure/database"
	"github.com/henrique502/go-repo-seed/infrastructure/external/opsgenie"
	_ "github.com/joho/godotenv/autoload"
)

var (
	DB       database.DatabaseClient
	Opsgenie opsgenie.OpsgenieClient
	AWS      aws.Config
)

func init() {
	var err error
	DB = database.New(os.Getenv("POSTGRE_URL"))
	Opsgenie = opsgenie.New()

	AWS, err = config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

}

func Close() {
	DB.Close()
}
