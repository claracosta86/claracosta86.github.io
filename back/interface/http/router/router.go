package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"poc2/back/infrastructure/container"
	"poc2/back/lib/logging"
)

func SetupRoutes(container *container.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Use(enableCors, logging.LoggingMiddleware)

	// Use handlers from the container
	commentHandler := container.CommentHandler
	culturalHandler := container.CulturalHandler
	notificationHandler := container.NotificationHandler
	userHandler := container.UserHandler

	// Rotas de comentários
	r.Route("/comments", func(r chi.Router) {
		r.Post("/", commentHandler.HandleCreateComment)
		r.Get("/{culturalType}/{culturalID:[0-9]+}", commentHandler.HandleGetComment)
	})

	// Rotas de culturais
	r.Route("/culturais", func(r chi.Router) {
		r.Get("/", culturalHandler.HandleGetAllCulturais)
		r.Get("/home", culturalHandler.HandleGetHomeCulturais)
		r.Post("/", culturalHandler.HandleCreateCultural)
		r.Get("/{type}/{id:[0-9]+}", culturalHandler.HandleGetCultural)
		r.Patch("/", culturalHandler.HandleUpdateCultural)
		r.Delete("/{type}/{id:[0-9]+}", culturalHandler.HandleDeleteCultural)
	})

	// Rotas de notificações
	r.Route("/notifications", func(r chi.Router) {
		r.Get("/{userID:[0-9]+}", notificationHandler.HandleGetUserNotifications)
		r.Patch("/{userID:[0-9]+}/seen", notificationHandler.HandleMarkNotificationsAsSeen)
	})

	// Rotas de usuário
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.HandleRegisterUser)
		r.Post("/login", userHandler.HandleUserLogin)
		r.Route("/{userID:[0-9]+}", func(r chi.Router) {
			r.Route("/profile", func(r chi.Router) {
				r.Get("/", userHandler.HandleGetUserProfile)
				r.Patch("/edit", userHandler.HandleEditUserProfile)
				r.Patch("/change-password", userHandler.HandleChangeUserPassword)
				r.Delete("/delete", userHandler.HandleDeleteUser)
			})
			r.Route("/favorites", func(r chi.Router) {
				r.Patch("/", userHandler.HandleFavorites)
				r.Get("/", userHandler.HandleGetUserFavorites)
				r.Patch("/last-seen", userHandler.HandleLastSeenFavorite)
			})
			r.Route("/culturais", func(r chi.Router) {
				r.Get("/", userHandler.HandleGetOrganizerCulturais)
				r.Get("/organizer", userHandler.HandleGetOrganizerInfo)
			})
		})
	})

	// Servidor de arquivos para assets estáticos
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	return r
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
