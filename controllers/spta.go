package controllers

import (
	"belajar_native/configs"
	"belajar_native/dtos"
	"belajar_native/helpers"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDataSpta(w http.ResponseWriter, r *http.Request) {
	db, err := configs.ConnectDB()
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	collection := db.Database("mitra_tani").Collection("spta")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	var results []dtos.SptaResponse
	if err := cursor.All(context.Background(), &results); err != nil {
		helpers.ResponseJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	helpers.ResponseJSON(w, http.StatusOK, results, "Success")
}
