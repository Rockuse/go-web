package main

import (
	"go-web/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web on port 8080")
	http.Handle("/", routes.Router())
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
