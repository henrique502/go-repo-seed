package api

import (
	"log"
	"net/http"
)

// Serve server http server
func Serve() {
	http.HandleFunc("/sync/alerts", HandlerFetchAlerts)
	http.HandleFunc("/sync/teams", HandlerFetchTeams)
	http.HandleFunc("/sync/integrations", HandlerFetchIntegrations)

	http.HandleFunc("/", HandlerFetchIntegrations)
	log.Println("Listing for requests at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
