package container

import (
	"database/sql"

	commentaryCase "poc2/back/application/comment"
	culturalCase "poc2/back/application/cultural"
	notificationCase "poc2/back/application/notification"
	userCase "poc2/back/application/user"
	commentaryService "poc2/back/domain/comment"
	culturalService "poc2/back/domain/cultural"
	notificationService "poc2/back/domain/notification"
	userService "poc2/back/domain/user"
	"poc2/back/infrastructure/persistence/mysql"
	http "poc2/back/interface/http/handlers"
)

// Container holds all the dependencies for the application
type Container struct {
	// Domain services
	CommentService      commentaryService.Service
	CulturalService     culturalService.Service
	NotificationService notificationService.Service
	UserService         userService.Service

	// Application use cases
	CommentUseCase      commentaryCase.UseCase
	CulturalUseCase     culturalCase.UseCase
	NotificationUseCase notificationCase.UseCase
	UserUseCase         userCase.UseCase

	// HTTP handlers
	CommentHandler      *http.CommentHandler
	CulturalHandler     *http.CulturalHandler
	NotificationHandler *http.NotificationHandler
	UserHandler         *http.UserHandler
}

// NewContainer creates a new dependency injection container
func NewContainer(db *sql.DB) *Container {
	// Infrastructure layer - repositories
	commentaryRepository := mysql.NewCommentRepository(db)
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
	commentaryHandler := http.NewCommentHandler(commentaryUseCase)
	culturalHandler := http.NewCulturalHandler(culturalUseCase)
	notificationHandler := http.NewNotificationHandler(notificationUseCase)
	userHandler := http.NewUserHandler(userUseCase)

	return &Container{
		CommentService:      commentaryService,
		CulturalService:     culturalService,
		UserService:         userService,
		NotificationService: notificationService,
		CommentUseCase:      commentaryUseCase,
		CulturalUseCase:     culturalUseCase,
		NotificationUseCase: notificationUseCase,
		UserUseCase:         userUseCase,
		CommentHandler:      commentaryHandler,
		CulturalHandler:     culturalHandler,
		NotificationHandler: notificationHandler,
		UserHandler:         userHandler,
	}
}
