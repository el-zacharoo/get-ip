package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/el-zacharoo/get-ip/auth"
	"github.com/el-zacharoo/get-ip/handler"
	"github.com/el-zacharoo/get-ip/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

var tokenAuth *auth.JWTAuth

func init() {
	jwks, _ := auth.JKS("https://practice-tenant.au.auth0.com/.well-known/jwks.json")
	tokenAuth = auth.New("RS256", jwks)
}

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

	r.Route("/post", func(r chi.Router) {
		r.Post("/", g.Create)
	})

	r.Route("/geo", func(r chi.Router) {
		// r.Post("/", g.Create)
		r.Use(
			auth.Verifier(tokenAuth),
			auth.Authenticator,
		)
		r.With(auth.Authz("read:entry")).Get("/{id}", g.Get)
		r.With(auth.Authz("read:entry")).Get("/", g.Query)

	})
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		fmt.Print(err)
	}
}
