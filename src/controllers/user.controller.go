package controllers

import (
	"fmt"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get All Users")
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get User By ID")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create User")
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update User By ID")
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete User By ID")
}
