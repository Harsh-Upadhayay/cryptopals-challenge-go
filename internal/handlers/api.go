package handlers

import (
	"cryptopals-challenge-go/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Capitalizaton: Function can be imported from other packages.
func Handler(r *chi.Mux) {

	// Global middleware
	// r struct will use stripSlashes middleware
	r.Use(chimiddle.StripSlashes)

	// r struct will be routed to this path,

	// Creates a subrouter specific to /account to group related routes and middleware.
	// Similar to django App's (incude).
	r.Route("/account", func(router chi.Router) {

		// Middleware specific to this path only /account/*
		// If the authorization fails then the response will be returned from here directly,
		// Without executing the rest of the code.
		router.Use(middleware.Authorization)

		// The passed function defines what'll happen at that path	GET: /account/counts/(django View)
		router.Get("/counts", GetCoinBalance)

	})

}
