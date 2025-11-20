package container

import (
	"database/sql"

	commentCase "poc2/back/application/comment"
	culturalCase "poc2/back/application/cultural"
	notificationCase "poc2/back/application/notification"
	userCase "poc2/back/application/user"
	commentService "poc2/back/domain/comment"
	culturalService "poc2/back/domain/cultural"
	notificationService "poc2/back/domain/notification"
	userService "poc2/back/domain/user"
	"poc2/back/infrastructure/persistence/mysql"
	http "poc2/back/interface/http/handlers"
)

// Container holds all the dependencies for the application
type Container struct {
	// Domain services
	CommentService      commentService.Service
	CulturalService     culturalService.Service
	NotificationService notificationService.Service
	UserService         userService.Service

	// Application use cases
	CommentUseCase      commentCase.UseCase
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
	commentRepository := mysql.NewCommentRepository(db)
	culturalRepository := mysql.NewCulturalRepository(db)
	notificationRepository := mysql.NewNotificationRepository(db)
	userRepository := mysql.NewUserRepository(db)

	// Domain layer - services
	commentService := commentService.NewService(commentRepository)
	culturalService := culturalService.NewService(culturalRepository)
	notificationService := notificationService.NewService(notificationRepository)
	userService := userService.NewService(userRepository)

	// Application layer - use cases
	commentUseCase := commentCase.NewUseCase(commentService, culturalService)
	culturalUseCase := culturalCase.NewUseCase(culturalService, userService)
	notificationUseCase := notificationCase.NewUseCase(notificationService, userService)
	userUseCase := userCase.NewUseCase(userService, culturalService)

	// Interface layer - HTTP handlers
	commentHandler := http.NewCommentHandler(commentUseCase)
	culturalHandler := http.NewCulturalHandler(culturalUseCase)
	notificationHandler := http.NewNotificationHandler(notificationUseCase)
	userHandler := http.NewUserHandler(userUseCase)

	return &Container{
		CommentService:      commentService,
		CulturalService:     culturalService,
		UserService:         userService,
		NotificationService: notificationService,
		CommentUseCase:      commentUseCase,
		CulturalUseCase:     culturalUseCase,
		NotificationUseCase: notificationUseCase,
		UserUseCase:         userUseCase,
		CommentHandler:      commentHandler,
		CulturalHandler:     culturalHandler,
		NotificationHandler: notificationHandler,
		UserHandler:         userHandler,
	}
}
