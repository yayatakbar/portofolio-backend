package public

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler untuk mendapatkan data diri dari MongoDB (public view)
func PersonalInfoHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := client.Database("portfolio").Collection("personal_info")
		ctx := context.TODO()

		var personalInfo models.PersonalInfo
		err := collection.FindOne(ctx, bson.M{}).Decode(&personalInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(personalInfo)
	}
}
