package server

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func StartServer(port int) {
	log.Println("Starting server on port :", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
