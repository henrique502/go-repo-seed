package main

import (
	"log"

	"github.com/henrique502/opsgenie/cmd/opsgenie"
	"github.com/henrique502/opsgenie/internal/integrations/jsonplaceholder"
	"github.com/joho/godotenv"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tracer.Start()
	jsonplaceholder.Init()
	opsgenie.Serve()
	defer tracer.Stop()
}
