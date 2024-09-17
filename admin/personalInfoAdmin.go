package admin

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio-backend/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func AddPersonalInfoHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var personalInfo models.PersonalInfo
		err := json.NewDecoder(r.Body).Decode(&personalInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := client.Database("portfolio").Collection("personal_info")
		ctx := context.TODO()

		_, err = collection.InsertOne(ctx, personalInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(personalInfo)
	}
}
