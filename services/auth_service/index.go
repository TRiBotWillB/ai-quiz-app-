package main

import (
	"api-test/infrastructure/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"log"
	"net/http"
)

func main() {
	auth.InitSupertokens()

	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	// SuperTokens Middleware
	r.Use(supertokens.Middleware)

	r.Get("/ping", session.VerifySession(nil, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}))
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
