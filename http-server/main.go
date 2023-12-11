package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.Host, r.Body, r.Header, r.Method)

		http.ServeFile(w, r, "./static/index.html")

	})

	log.Fatal(http.ListenAndServe(":50052", nil))

}
