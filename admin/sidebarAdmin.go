package admin

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio-backend/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddSidebarHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sidebar models.Sidebar
		err := json.NewDecoder(r.Body).Decode(&sidebar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := client.Database("portfolio").Collection("sidebar")
		ctx := context.TODO()

		_, err = collection.InsertOne(ctx, sidebar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(sidebar)
	}
}
