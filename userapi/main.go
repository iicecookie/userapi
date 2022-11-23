package main

import (
	"net/http"
	"refactoring/controller"
	"refactoring/repository"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	// Вынести в Di
	var jsonUserRepository repository.Repository = repository.NewUserJsonRepository()
	userController := controller.NewUserController(jsonUserRepository)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", userController.SearchUsers)
				r.Post("/", userController.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", userController.GetUser)
					r.Patch("/", userController.UpdateUser)
					r.Delete("/", userController.DeleteUser)
				})
			})
		})
	})

	http.ListenAndServe(":3333", r)
}
