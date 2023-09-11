package main

import (
	"fmt"
	"inventory/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", handlers.AddItemHandler)
	http.HandleFunc("/view", handlers.ViewItemsHandler)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
