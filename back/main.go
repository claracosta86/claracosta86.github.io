package main

import (
    "log"
    "net/http"

    "poc2/back/routes"
    "poc2/back/repository"
    "poc2/back/service"

)


func main() {
    log.Println("Servidor rodando em http://localhost:8080")

	userRepo := repository.NewUserRepository()
	eventRepo := repository.NewEventRepository()
	attractionRepo := repository.NewAttractionRepository()

	userService := service.NewUserService(userRepo, eventRepo, attractionRepo)
	eventService := service.NewEventService()
	attractionService := service.NewAttractionService()

	router := routes.SetupRoutes(userService, eventService, attractionService)

	http.ListenAndServe(":8080", router)
}