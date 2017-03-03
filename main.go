package main

import (
	"log"
	"net/http"

	"github.com/deep/GOMVC/models"
)

func main() {

	router := Router()
	models.InitializeGoDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}
