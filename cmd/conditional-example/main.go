package main

import (
	"net/http"

	"github.com/a-h/templ"

	"example-templ-project/internal/views"
)

func main() {
	http.Handle("/", templ.Handler(views.CatContent("Cat Page!")))
	http.HandleFunc("/conditional", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// Full page with layout for initial load
			views.CatTreatingPage("Whiskers", false).Render(r.Context(), w)
		} else {
			// Just the content for HTMX requests
			hasTreats := r.URL.Query().Get("hasTreats") == "true"
			views.CatTreatingContent("Whiskers", hasTreats).Render(r.Context(), w)
		}
	})

	http.ListenAndServe(":8080", nil)
}
