package main

import (
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/henrique502/go-repo-seed/domain/alerts"
	_ "github.com/henrique502/go-repo-seed/infra/opsgenie"
	_ "github.com/henrique502/go-repo-seed/infra/postgre"
)

func main() {
	// DataDog
	// tracer.Start()
	// defer tracer.Stop()

	// API
	// api.Serve()

	alerts.FetchDay(time.Now())
}
