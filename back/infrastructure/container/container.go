package container

import (
	"database/sql"

	userCase "poc2/back/application/user"
	userService "poc2/back/domain/user"
	"poc2/back/infrastructure/persistence/mysql"
	"poc2/back/interface/http"

)

// Container holds all the dependencies for the application
type Container struct {
	// Domain services
	UserService userService.Service
	
	// Application use cases
	UserUseCase userCase.UseCase

	// HTTP handlers
	UserHandler *http.UserHandler
}

// NewContainer creates a new dependency injection container
func NewContainer(db *sql.DB) *Container {
	// Infrastructure layer - repositories
	userRepository := mysql.NewUserRepository(db)
	
	// Domain layer - services
	userService := userService.NewService(userRepository)
	
	// Application layer - use cases
	userUseCase := userCase.NewUseCase(userService)

	// Interface layer - HTTP handlers
	userHandler := http.NewUserHandler(userUseCase)
	
	return &Container{
		UserService: userService,
		UserUseCase: userUseCase,
		UserHandler: userHandler,
	}
}
