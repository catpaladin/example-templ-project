package main

import (
	"net/http"
	"strconv"

	"example-templ-project/internal/views"
)

var (
	// This is mock data. Generally you would get and update this data with a database
	cats = []views.Cat{
		{ID: 1, Name: "Finn", Breed: "Orange", Naps: 5, FavToy: "Mouse"},
		{ID: 2, Name: "Honey", Breed: "Calico", Naps: 3, FavToy: "String"},
		{ID: 3, Name: "Wonkers", Breed: "Black", Naps: 7, FavToy: "Catnip Ball"},
	}
	globalHasTreats = false
)

func main() {
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/pet-cat/", handlePetCat)
	http.HandleFunc("/treats", handleTreats)

	http.ListenAndServe(":8080", nil)
}

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	views.CatPage(cats, globalHasTreats).Render(r.Context(), w)
}

func handlePetCat(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/pet-cat/"):]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid cat ID", http.StatusBadRequest)
		return
	}

	var targetCat views.Cat
	var found bool

	for i, cat := range cats {
		if cat.ID == id {
			cats[i].Naps++
			targetCat = cats[i]
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Cat not found", http.StatusNotFound)
		return
	}

	views.CatListItem(targetCat).Render(r.Context(), w)
}

func handleTreats(w http.ResponseWriter, r *http.Request) {
	// Toggle treats state
	hasTreatsParam := r.URL.Query().Get("hasTreats")
	globalHasTreats = (hasTreatsParam == "true")

	views.CatTreatingSection(globalHasTreats).Render(r.Context(), w)
}
