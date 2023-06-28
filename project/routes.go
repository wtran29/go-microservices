package main

import (
	"net/http"

	"github.com/wtran29/fenix/fenix"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes

	// uncomment if implementing 'remember me' login
	// a.use(a.Middleware.CheckRemember)

	// add routes here
	a.get("/", a.Handlers.Home)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from fenix
	a.App.Routes.Mount("/fenix", fenix.Routes())
	a.App.Routes.Mount("/api", a.ApiRoutes())

	return a.App.Routes
}
