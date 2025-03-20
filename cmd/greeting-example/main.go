package main

import (
	"net/http"

	"github.com/a-h/templ"

	"example-templ-project/internal/views"
)

func main() {
	http.Handle("/", templ.Handler(views.CatGreeting("Whiskers")))

	http.ListenAndServe(":8080", nil)
}
