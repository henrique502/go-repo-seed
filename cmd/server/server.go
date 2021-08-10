package server

import (
	"io"
	"log"
	"net/http"

	"github.com/henrique502/go-repo-seed/internal/integrations/jsonplaceholder"
)

// Serve server http server
func Serve() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, jsonplaceholder.ListPosts())
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:3000/hello")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
