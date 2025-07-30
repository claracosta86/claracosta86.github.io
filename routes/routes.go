package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"poc2/back/api"
	"poc2/back/service"
	"poc2/back/lib/logging"
	"poc2/front/handlers"

)

func SetupRoutes(userService service.UserService, eventService service.EventService, attractionService service.AttractionService) *chi.Mux {
	r := chi.NewRouter()

	r.Use(enableCors, logging.LoggingMiddleware)

	userHandler := api.NewUserHandler(userService, eventService, attractionService)
	eventHandler := api.NewEventHandler()
	attractionHandler := api.NewAttractionHandler()

	// Rotas de usuário
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.HandleRegisterUser)
		r.Get("/fetch", userHandler.HandleGetUserData)
		r.Route("/favorites", func (r chi.Router) {
			r.Get("/{id}", userHandler.HandleGetUserFavorites)
			r.Route("/{id}/event", func (r chi.Router) {
				r.Post("/add", userHandler.HandleAddEventToFavorites)
				r.Delete("/delete", userHandler.HandleDeleteEventFromFavorites)
			})
			r.Route("/{id}/attraction", func (r chi.Router) {
				r.Post("/add", userHandler.HandleAddAttractionToFavorites)
				r.Delete("/delete", userHandler.HandleDeleteAttractionFromFavorites)
			})
		})
		r.Put("/update", userHandler.HandleUpdateUser)
		r.Delete("/delete/{id}", userHandler.HandleDeleteUser)

	})

	// Rotas de eventos
	r.Route("/events", func(r chi.Router) {
		r.Post("/register", eventHandler.HandleRegisterEvent)
		r.Get("/all", eventHandler.HandleGetAllEvents)
		r.Get("/{id}", eventHandler.HandleGetEventByID)
		r.Put("/update", eventHandler.HandleUpdateEvent)
		r.Delete("/delete/{id}", eventHandler.HandleDeleteEvent)
	})

	// Rotas de atrações
	r.Route("/attractions", func(r chi.Router) {
		r.Post("/register", attractionHandler.HandleRegisterAttraction)
		r.Get("/all", attractionHandler.HandleGetAllAttractions)
		r.Get("/{id}", attractionHandler.HandleGetAttractionByID)
		r.Put("/update", attractionHandler.HandleUpdateAttraction)
		r.Delete("/delete/{id}", attractionHandler.HandleDeleteAttraction)
	})


	templatesHandler := handlers.NewTemplatesHandler()

	// Rotas dos templates
	fs := http.FileServer(http.Dir("./docs"))
	r.Handle("/*", fs)
	
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/index.html")
	})

	r.Post("/role/select", templatesHandler.HandleUserTypeSelection)
	r.Get("/login/", templatesHandler.HandleLogin)      // login único
	r.Get("/register/", templatesHandler.HandleRegistry) // registro único

	return r
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}