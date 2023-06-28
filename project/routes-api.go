package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) ApiRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(mux chi.Router) {
		// add any api routes here
		r.Get("/test-api", func(w http.ResponseWriter, r *http.Request) {
			var payload struct {
				Content string `json:"content"`
			}

			payload.Content = "hello world"
			a.App.WriteJSON(w, http.StatusOK, payload)
		})
	})

	return r
}
