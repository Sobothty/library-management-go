package main

import (
	"Go-Project/feature/Config"
	"Go-Project/feature/Controller"
	"Go-Project/feature/Repository"
	"Go-Project/feature/Routes"
	"Go-Project/feature/Service"
	"log"
	"net/http"
)

func main() {
	Config.ConnectDB()

	repository := Repository.ConnectRepository(Config.ConnectDB())
	service := Service.NewBookServiceImpl(repository)
	controller := Controller.NewBookController(service)

	router := Routes.NewRoutes(controller)

	mux := http.NewServeMux()
	router.RoutesEndpoint(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
