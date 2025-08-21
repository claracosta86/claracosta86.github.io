package container

import (
	"database/sql"

	userCase "poc2/back/application/user"
	userService "poc2/back/domain/user"
	notificationCase "poc2/back/application/notification"
	notificationService "poc2/back/domain/notification"
	"poc2/back/infrastructure/persistence/mysql"
	http "poc2/back/interface/http/handlers"
)

// Container holds all the dependencies for the application
type Container struct {
	// Domain services
	UserService userService.Service
	// CulturalService culturalService.Service
	NotificationService notificationService.Service

	// Application use cases
	UserUseCase userCase.UseCase
	// CulturalUseCase culturalCase.UseCase
	NotificationUseCase notificationCase.UseCase

	// HTTP handlers
	UserHandler *http.UserHandler
	// CulturalHandler *http.CulturalHandler
	NotificationHandler *http.NotificationHandler
}

// NewContainer creates a new dependency injection container
func NewContainer(db *sql.DB) *Container {
	// Infrastructure layer - repositories
	userRepository := mysql.NewUserRepository(db)
	notificationRepository := mysql.NewNotificationRepository(db)
	
	// Domain layer - services
	userService := userService.NewService(userRepository)
	notificationService := notificationService.NewService(notificationRepository)

	// Application layer - use cases
	userUseCase := userCase.NewUseCase(userService)
	notificationUseCase := notificationCase.NewUseCase(notificationService)

	// Interface layer - HTTP handlers
	userHandler := http.NewUserHandler(userUseCase)
	notificationHandler := http.NewNotificationHandler(notificationUseCase)

	return &Container{
		UserService:      userService,
		NotificationService: notificationService,
		UserUseCase:      userUseCase,
		NotificationUseCase: notificationUseCase,
		UserHandler:      userHandler,
		NotificationHandler: notificationHandler,
	}
}
