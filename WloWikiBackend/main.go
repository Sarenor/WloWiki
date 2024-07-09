// main.go
package main

import (
	"WloWikiBackend/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Item represents a simple item structure
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// In-memory database (for demonstration purposes)
var items []Item
var nextID = 1

func main() {
	r := mux.NewRouter()
	client, err := mongo.NewClient("mongodb://admin:<PASS>@49.13.132.251:27017")
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Create the Items collection instance
	itemsCollection := mongo.NewItemsCollection(client)

	// // Define your items
	// items := []mongo.Item{
	// 	{Image: "src/routes/items/images/silvery_star.png", Rank: "37 (76)", Name: "Silvery Star", Type: "Armor (Head)", Base: "Corundum / Platinum", Attributes: "MAT: +37"},
	// 	{Image: "src/routes/items/images/voodoo_doll.png", Rank: "32 (64)", Name: "Voodo Doll", Type: "Armor (Head)", Base: "Corundum / Platinum", Attributes: "MAT: +37"},
	// 	{Image: "src/routes/items/images/jade_helmet.png", Rank: "26 (52)", Name: "Jade Helmet", Type: "Armor (Head)", Base: "Crystallization / Titanium / Magic", Attributes: "SPD: +26"},
	// }

	// // Insert Items
	// if err := itemsCollection.InsertMany(items); err != nil {
	// 	log.Fatalf("Failed to insert documents: %v", err)
	// }

	//fmt.Println("Items inserted successfully!")

	// Fetch and display items
	// fetchedItems, err := itemsCollection.FindAll()
	// if err != nil {
	// 	log.Fatalf("Failed to fetch documents: %v", err)
	// }

	// for _, item := range fetchedItems {
	// 	fmt.Println(item)
	// }

	r.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			// Preflight request
			return
		}
		switch r.Method {
		case http.MethodGet:
			fetchItemsHandler(w, r, itemsCollection)
		case http.MethodPost:
			createItemHandler(w, r, itemsCollection)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// fetchItemsHandler handles the GET request to fetch all items
func fetchItemsHandler(w http.ResponseWriter, r *http.Request, collection *mongo.ItemsCollection) {
	items, err := collection.FindAll()
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		log.Printf("Error fetching items: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, "Failed to encode items", http.StatusInternalServerError)
		log.Printf("Error encoding items: %v", err)
	}
}

// createItemHandler handles the POST request to create a new item
func createItemHandler(w http.ResponseWriter, r *http.Request, itemsCollection *mongo.ItemsCollection) {
	fmt.Print(r)
	var item mongo.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := itemsCollection.InsertOne(item)
	if err != nil {
		http.Error(w, "Failed to add item", http.StatusInternalServerError)
		log.Printf("Error adding item: %v", err)
		return
	}

	response := map[string]string{"message": "Item added successfully"}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		log.Printf("Error marshalling response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
