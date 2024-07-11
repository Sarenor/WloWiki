// main.go
package main

import (
	"WloWikiBackend/mongo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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
	client, err := mongo.NewClient("mongodb://admin:ojNBLAZXEcqc2@49.13.132.251:27017")
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Create the Items collection instance
	itemsCollection := mongo.NewItemsCollection(client)

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

	// New endpoint for fetching a single item by item_id
	r.HandleFunc("/api/items/{item_id}", func(w http.ResponseWriter, r *http.Request) {
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
			fetchItemHandler(w, r, itemsCollection)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func fetchItemsHandler(w http.ResponseWriter, r *http.Request, collection *mongo.ItemsCollection) {
	pageNumber, _ := strconv.Atoi(r.URL.Query().Get("_page"))
	rowsPerPage, _ := strconv.Atoi(r.URL.Query().Get("_limit"))
	sortOrder := r.URL.Query().Get("_order")
	sortBy := r.URL.Query().Get("_sort")
	searchQuery := r.URL.Query().Get("q")

	// Default values
	if pageNumber == 0 {
		pageNumber = 1
	}
	if rowsPerPage == 0 {
		rowsPerPage = 10
	}

	// Build the filter
	filter := bson.M{}
	if searchQuery != "" {
		filter["name"] = bson.M{"$regex": searchQuery, "$options": "i"}
	}

	// Define sort options
	var sortOptions bson.D
	if sortBy != "" {
		order := 1
		if sortOrder == "desc" {
			order = -1
		}
		sortOptions = bson.D{{sortBy, order}}
	} else {
		sortOptions = bson.D{{"_id", 1}} // Default sort by _id ascending
	}

	// Define pagination parameters
	skip := int64((pageNumber - 1) * rowsPerPage)
	limit := int64(rowsPerPage)

	// Fetch the items
	items, err := collection.FindAll(filter, sortOptions, skip, limit)
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

func fetchItemHandler(w http.ResponseWriter, r *http.Request, collection *mongo.ItemsCollection) {
	vars := mux.Vars(r)
	itemID := vars["item_id"]

	// Convert itemID to MongoDB ObjectID
	objID, err := strconv.Atoi(itemID)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		log.Printf("Invalid item ID: %v", err)
		return
	}

	// Build the filter
	filter := bson.M{"item_id": objID}

	// Fetch the item
	item, err := collection.FindOne(filter)
	if err != nil {
		http.Error(w, "Failed to fetch item", http.StatusInternalServerError)
		log.Printf("Error fetching item: %v", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, "Failed to encode item", http.StatusInternalServerError)
		log.Printf("Error encoding item: %v", err)
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
