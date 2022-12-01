package routes

import (
	"github.com/gorilla/mux"

	"golang-api/src/controllers"
)

func InitUserRoutes(api *mux.Router) {
	user := api.PathPrefix("/user").Subrouter()

	user.HandleFunc("", controllers.GetAllUsers).Methods("GET")
}
