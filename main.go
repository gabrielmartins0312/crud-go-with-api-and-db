package main

import (
	"log"
	"net/http"

	"github.com/gabrielmartins0312/crud-go-with-api-and-db/config"
	"github.com/gabrielmartins0312/crud-go-with-api-and-db/middleware"
	"github.com/gabrielmartins0312/crud-go-with-api-and-db/router"
)

func main() {
	config.Connect()

	mux := http.NewServeMux()
	router.LoadRoutes(mux)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.EnableCORS(mux)))
}
