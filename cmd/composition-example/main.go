package main

import (
	"net/http"

	"github.com/a-h/templ"

	"example-templ-project/internal/views"
)

func main() {
	http.Handle("/", templ.Handler(views.CatContent("Cat Page!")))

	http.ListenAndServe(":8080", nil)
}
