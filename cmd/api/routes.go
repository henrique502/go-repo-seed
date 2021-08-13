package api

import (
	"net/http"

	"github.com/henrique502/go-repo-seed/domain/alerts"
	"github.com/henrique502/go-repo-seed/domain/integrations"
	"github.com/henrique502/go-repo-seed/domain/teams"
)

func HandlerFetchAlerts(w http.ResponseWriter, req *http.Request) {
	alerts.Sync()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func HandlerFetchTeams(w http.ResponseWriter, req *http.Request) {
	teams.Sync()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func HandlerFetchIntegrations(w http.ResponseWriter, req *http.Request) {
	integrations.Sync()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
