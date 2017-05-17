package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
