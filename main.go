package main

import (
	"log"

	"github.com/henrique502/go-repo-seed/cmd/server"
	"github.com/henrique502/go-repo-seed/internal/integrations/jsonplaceholder"
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
	server.Serve()
	defer tracer.Stop()
}
