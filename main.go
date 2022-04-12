package main

import (
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
			AllowedMethods:   []string{"GET", "POST"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
			// Debug:            true,
		}),
	)

	add := &handler.Geolocation{
		Store: s,
	}

	r.Get("/ip", handler.Get)
	r.Post("/geo", add.Add)

	http.ListenAndServe(":8080", r)
}
