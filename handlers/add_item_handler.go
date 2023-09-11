package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// GroceryItem represents a grocery item.
type GroceryItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var inventory []GroceryItem
var mu sync.Mutex

func init() {
	loadInventory()
}

func loadInventory() {
	file, err := os.OpenFile("data/inventory.json", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Error opening inventory file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&inventory); err != nil {
		log.Println("Error decoding inventory:", err)
	}
}

func saveInventory() {
	file, err := os.OpenFile("data/inventory.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Println("Error opening inventory file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(inventory); err != nil {
		log.Println("Error encoding inventory:", err)
	}
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.FormValue("name")
		priceStr := r.FormValue("price") // Get price as a string

		// Convert priceStr to a float64
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Invalid price format:", err)
			http.Error(w, "Invalid price format", http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()

		newItem := GroceryItem{
			ID:    len(inventory) + 1,
			Name:  name,
			Price: price, // Assign the float64 price
		}

		inventory = append(inventory, newItem)
		saveInventory()

		log.Printf("Added item: ID=%d, Name=%s, Price=%.2f", newItem.ID, newItem.Name, newItem.Price)

		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/add_item.html")
	if err != nil {
		log.Println("Error parsing add_item.html template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
