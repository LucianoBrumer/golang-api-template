package controllers

import (
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get All Users")
}

func CreateUser(name string, lastname string) {
	// db.Database()
	// db.DB
}
