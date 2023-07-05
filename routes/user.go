package routes

import (
	"belajar_native/controllers"

	"github.com/gorilla/mux"
)

func Route() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/api/v1/spta", controllers.GetDataSpta).Methods("GET")

	return router
}
