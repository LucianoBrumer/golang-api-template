package routes

import (
	"github.com/gorilla/mux"

	"golang-api/src/controllers"
)

func InitUserRoutes(api *mux.Router) {
	user := api.PathPrefix("/user").Subrouter()

	user.HandleFunc("", controllers.GetAllUsers).Methods("GET")
	user.HandleFunc("/{id}", controllers.GetUserByID).Methods("GET")
	user.HandleFunc("", controllers.CreateUser).Methods("POST")
	user.HandleFunc("/{id}", controllers.UpdateUserByID).Methods("GET")
	user.HandleFunc("/{id}", controllers.DeleteUserByID).Methods("DELETE")
}
