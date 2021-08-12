package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/henrique502/go-repo-seed/infra/opsgenie"
)

// Serve server http server
func Serve() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		alerts := opsgenie.GetAlertList(time.Now())
		err := json.NewEncoder(w).Encode(alerts)
		if err != nil {
			log.Panic(err)
		}
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:3000/hello")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
