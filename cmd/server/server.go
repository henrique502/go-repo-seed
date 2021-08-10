package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/henrique502/go-repo-seed/internal/integrations/jsonplaceholder"
)

// Serve server http server
func Serve() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		posts := jsonplaceholder.ListPosts()
		err := json.NewEncoder(w).Encode(posts)
		if err != nil {
			log.Fatalln(err)
		}
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:3000/hello")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
