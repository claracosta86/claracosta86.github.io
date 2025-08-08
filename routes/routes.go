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

	// Handlers Backend
	
	userHandler := api.NewUserHandler(userService, eventService, attractionService)
	eventHandler := api.NewEventHandler(eventService)
	attractionHandler := api.NewAttractionHandler(attractionService)

	// Rotas de usuário
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.HandleRegisterUser)
		r.Post("/login", userHandler.HandleUserLogin)
		r.Route("/{id}", func(r chi.Router) {
			r.Route("/profile", func (r chi.Router) {
				r.Get("/", userHandler.HandleGetUserProfile)
				r.Put("/edit", userHandler.HandleEditUserProfile)
				r.Put("/change-password", userHandler.HandleChangeUserPassword)
			})
			r.Route("/favorites", func (r chi.Router) {
				r.Get("/", userHandler.HandleGetUserFavorites)
				r.Route("/{type}/{typeID}", func (r chi.Router) {
					r.Post("/add", userHandler.HandleAddToFavorites)
					r.Delete("/delete", userHandler.HandleDeleteFromFavorites)
				})
			})
		})
		r.Delete("/delete/{id}", userHandler.HandleDeleteUser)
	})

	// Rotas de eventos e atrações
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


	// Handlers Frontend

	templatesHandler := handlers.NewTemplatesHandler()

	// Rotas dos templates
	r.Handle("/loginpage/*",
	http.StripPrefix("/loginpage/",
		http.FileServer(http.Dir("./docs/loginpage"))))

	r.Handle("/registerpage/*",
	http.StripPrefix("/registerpage/",
		http.FileServer(http.Dir("./docs/registerpage"))))

	r.Handle("/homepage/*",
	http.StripPrefix("/homepage/",
		http.FileServer(http.Dir("./docs/homepage"))))

	r.Handle("/images/*",
	http.StripPrefix("/images/",
		http.FileServer(http.Dir("./docs/images"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    	http.ServeFile(w, r, "./docs/index.html")
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/select-type", templatesHandler.HandleUserTypeSelection)
		r.Get("/login", templatesHandler.HandleLogin)
		r.Get("/register/", templatesHandler.HandleRegistry)
		r.Get("/password-recovery", templatesHandler.HandleForgottenPassword)
		r.Route("/profile", func(r chi.Router) {
			r.Get("/", templatesHandler.HandleProfile)
			r.Get("/edit", templatesHandler.HandleEditProfile)
			r.Get("/change-password", templatesHandler.HandleChangePassword)
		})
		r.Get("/favorites", templatesHandler.HandleViewFavorites)
	})
	r.Get("/home", templatesHandler.HandleHome)

	// Fallback para renderizar arquivos estáticos
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("./docs"))))

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