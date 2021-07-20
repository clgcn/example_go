package main

import (
	"log"
	"net/http"

	"github.com/charlesguo404/example_go/database"
	"github.com/charlesguo404/example_go/product"
	"github.com/charlesguo404/example_go/receipt"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
