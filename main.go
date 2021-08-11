package main

import (
	_ "github.com/joho/godotenv/autoload"

	"github.com/henrique502/go-repo-seed/cmd/api"
	_ "github.com/henrique502/go-repo-seed/infra/jsonplaceholder"
	_ "github.com/henrique502/go-repo-seed/infra/opsgenie"
	_ "github.com/henrique502/go-repo-seed/infra/postgre"
)

func main() {
	// DataDog
	// tracer.Start()
	// defer tracer.Stop()

	// API
	api.Serve()
}
