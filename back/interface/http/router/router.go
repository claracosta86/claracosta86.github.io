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
	culturalHandler := container.CulturalHandler
	notificationHandler := container.NotificationHandler
	userHandler := container.UserHandler

	// Rotas de usuário
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.HandleRegisterUser)
		r.Post("/login", userHandler.HandleUserLogin)
		r.Route("/{userID:[0-9]+}", func(r chi.Router) {
			r.Route("/profile", func (r chi.Router) {
				r.Get("/", userHandler.HandleGetUserProfile)
				r.Patch("/edit", userHandler.HandleEditUserProfile)
				r.Patch("/change-password", userHandler.HandleChangeUserPassword)
				r.Delete("/delete", userHandler.HandleDeleteUser)
			})
			r.Route("/favorites", func (r chi.Router) {
				r.Patch("/", userHandler.HandleFavorites)
				r.Get("/", userHandler.HandleGetUserFavorites)
				r.Patch("/last-seen", userHandler.HandleLastSeenFavorite)
			})
			r.Route("/culturais", func (r chi.Router) {
				r.Get("/", userHandler.HandleGetOrganizerCulturais)
				// r.Get("/organizer", userHandler.HandleGetOrganizerInformation)
			})
		})
		r.Post("/select-type", userHandler.HandleUserTypeSelection)
		r.Get("/get-type", userHandler.HandleGetUserType)
		r.Post("/set-information", userHandler.HandleSetUserInformation)
		r.Get("/get-information", userHandler.HandleGetUserInformation)
	})

	r.Route("/culturais", func(r chi.Router) {
		r.Post("/", culturalHandler.HandleCreateCultural)
		r.Get("/{type}/{id:[0-9]+}", culturalHandler.HandleGetCultural)
		r.Patch("/{type}/{id:[0-9]+}", culturalHandler.HandleUpdateCultural)
		r.Delete("/{type}/{id:[0-9]+}", culturalHandler.HandleDeleteCultural)
	})

	r.Route("/notifications", func(r chi.Router) {
		r.Get("/{userID:[0-9]+}", notificationHandler.HandleGetUserNotifications)
		r.Patch("/{userID:[0-9]+}/seen", notificationHandler.HandleMarkNotificationsAsSeen)
	})
	
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