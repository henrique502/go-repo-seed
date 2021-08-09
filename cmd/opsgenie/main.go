package opsgenie

import (
	"io"
	"log"
	"net/http"
)

func Serve() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:3000/hello")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
