package api

import (
	"fmt"
	"net/http"

	"github.com/henrique502/go-repo-seed/application/sync"
)

func HandlerFetchAlerts(w http.ResponseWriter, req *http.Request) {
	s := sync.New()
	err := s.Alerts()

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func HandlerFetchTeams(w http.ResponseWriter, req *http.Request) {
	s := sync.New()
	err := s.Teams()

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func HandlerFetchIntegrations(w http.ResponseWriter, req *http.Request) {
	s := sync.New()
	err := s.Integrations()

	if err != nil {
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
