package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

func ViewItemsHandler(w http.ResponseWriter, r *http.Request) {
	var mu sync.Mutex
	tmpl, err := template.ParseFiles("templates/view_items.html")
	if err != nil {
		log.Println("Error parsing view_items.html template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	tmpl.Execute(w, inventory)
}
