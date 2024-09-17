package public

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio-backend/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler untuk mendapatkan data diri dari MongoDB (public view)
func SidebarHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set timeout for MongoDB operations
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Access collection
		collection := client.Database("portfolio").Collection("sidebar")

		// Prepare a cursor to retrieve multiple documents
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, "Failed to retrieve sidebars: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)

		// Create a slice to hold the sidebar data
		var sidebars []models.Sidebar

		// Iterate through the cursor and decode each document
		for cursor.Next(ctx) {
			var sidebar models.Sidebar
			err := cursor.Decode(&sidebar)
			if err != nil {
				http.Error(w, "Error decoding sidebar: "+err.Error(), http.StatusInternalServerError)
				return
			}
			sidebars = append(sidebars, sidebar)
		}

		// Check for any cursor errors
		if err := cursor.Err(); err != nil {
			http.Error(w, "Cursor error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Encode the result as JSON and return to the client
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(sidebars)
		if err != nil {
			http.Error(w, "Failed to encode sidebar data: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
