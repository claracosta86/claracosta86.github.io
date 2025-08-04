package main

import (
    "log"
    "net/http"
	"database/sql"

    _ "github.com/go-sql-driver/mysql"

    "poc2/routes"
    "poc2/back/repository"
    "poc2/back/service"

)

var db *sql.DB

func main() {
    log.Println("Servidor rodando em http://localhost:8080/")

	dsn := "claracosta86:sua_senha@tcp(127.0.0.1:3306)/POCII"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verifica se está conectado
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository()
	attractionRepo := repository.NewAttractionRepository(db, userRepo)

	userService := service.NewUserService(userRepo, eventRepo, attractionRepo)
	eventService := service.NewEventService(eventRepo)
	attractionService := service.NewAttractionService(attractionRepo)

	router := routes.SetupRoutes(userService, eventService, attractionService)

	http.ListenAndServe(":8080", router)
}