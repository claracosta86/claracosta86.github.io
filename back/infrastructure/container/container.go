package container

import (
	"database/sql"

	commentaryCase "poc2/back/application/commentary"
	commentaryService "poc2/back/domain/commentary"
	culturalCase "poc2/back/application/cultural"
	culturalService "poc2/back/domain/cultural"
	notificationCase "poc2/back/application/notification"
	notificationService "poc2/back/domain/notification"
	userCase "poc2/back/application/user"
	userService "poc2/back/domain/user"
	"poc2/back/infrastructure/persistence/mysql"
	http "poc2/back/interface/http/handlers"
)

// Container holds all the dependencies for the application
type Container struct {
	// Domain services
	CommentaryService commentaryService.Service
	CulturalService culturalService.Service
	NotificationService notificationService.Service
	UserService userService.Service

	// Application use cases
	CommentaryUseCase commentaryCase.UseCase
	CulturalUseCase culturalCase.UseCase
	NotificationUseCase notificationCase.UseCase
	UserUseCase userCase.UseCase

	// HTTP handlers
	CommentaryHandler *http.CommentaryHandler
	CulturalHandler   *http.CulturalHandler
	NotificationHandler *http.NotificationHandler
	UserHandler       *http.UserHandler
}

// NewContainer creates a new dependency injection container
func NewContainer(db *sql.DB) *Container {
	// Infrastructure layer - repositories
	commentaryRepository := mysql.NewCommentaryRepository(db)
	culturalRepository := mysql.NewCulturalRepository(db)
	notificationRepository := mysql.NewNotificationRepository(db)
	userRepository := mysql.NewUserRepository(db)

	// Domain layer - services
	commentaryService := commentaryService.NewService(commentaryRepository)
	culturalService := culturalService.NewService(culturalRepository)
	notificationService := notificationService.NewService(notificationRepository)
	userService := userService.NewService(userRepository)

	// Application layer - use cases
	commentaryUseCase := commentaryCase.NewUseCase(commentaryService, culturalService)
	culturalUseCase := culturalCase.NewUseCase(culturalService, userService)
	notificationUseCase := notificationCase.NewUseCase(notificationService, userService)
	userUseCase := userCase.NewUseCase(userService, culturalService)

	// Interface layer - HTTP handlers
	commentaryHandler := http.NewCommentaryHandler(commentaryUseCase)
	culturalHandler := http.NewCulturalHandler(culturalUseCase)
	notificationHandler := http.NewNotificationHandler(notificationUseCase)
	userHandler := http.NewUserHandler(userUseCase)

	return &Container{
		CommentaryService: commentaryService,
		CulturalService:   culturalService,
		UserService:       userService,
		NotificationService: notificationService,
		CommentaryUseCase: commentaryUseCase,
		CulturalUseCase:   culturalUseCase,
		NotificationUseCase: notificationUseCase,
		UserUseCase:      userUseCase,
		CommentaryHandler: commentaryHandler,
		CulturalHandler:  culturalHandler,
		NotificationHandler: notificationHandler,
		UserHandler:      userHandler,
	}
}
