package main

import (
	"Go-Project/feature/Config"
	"Go-Project/feature/Controller"
	"Go-Project/feature/Repository"
	"Go-Project/feature/Routes"
	"Go-Project/feature/Service"
	"Go-Project/feature/middleware"
	"log"
	"net/http"
)

func main() {
	Config.ConnectDB()

	repository := Repository.ConnectRepository(Config.ConnectDB())
	service := Service.NewBookServiceImpl(repository)
	controller := Controller.NewBookController(service)

	router := Routes.NewRoutes(controller)

	// Start Server
	mux := http.NewServeMux()
	router.RoutesEndpoint(mux)

	//Add Middleware
	handler := middleware.LoggingMiddleware(middleware.JsonMiddleware(mux))

	log.Fatal(http.ListenAndServe(":8080", handler))
}
