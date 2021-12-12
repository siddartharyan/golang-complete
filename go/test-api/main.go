package main

import (
	"log"
	"net/http"
	"test-api/router"
)


func main() {
	router := router.Router()
	log.Fatal(http.ListenAndServe(":3000",router))
}