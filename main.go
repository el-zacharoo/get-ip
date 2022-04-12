package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ip-address/handler"
	"github.com/ip-address/store"
)

func main() {
	s := store.Connect()
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.StripSlashes,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "QUERY"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
			// Debug:            true,
		}),
	)

	g := &handler.Geolocation{
		Store: s,
	}

	r.Get("/ip", handler.Get)
	r.Route("/geo", func(r chi.Router) {
		r.Post("/", g.Create)
		// r.Get("/{id}", g.Get)
		// r.Put("/{id}", g.Update)
		// r.Delete("/{id}", g.Delete)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Print(err)
	}
}
