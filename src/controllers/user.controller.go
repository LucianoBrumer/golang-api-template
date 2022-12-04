package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"data": "GetAllUsers",
	}

	json.NewEncoder(w).Encode(data)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	data := map[string]interface{}{
		"data": "GetUserByID",
		"id":   vars["id"],
	}

	json.NewEncoder(w).Encode(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	byt := []byte(bodyData)
	var bodyJson map[string]string
	json.Unmarshal(byt, &bodyJson)

	data := map[string]interface{}{
		"data": "CreateUser",
		"body": bodyJson,
	}

	json.NewEncoder(w).Encode(data)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	data := map[string]interface{}{
		"data": "UpdateUserByID",
		"id":   vars["id"],
	}

	json.NewEncoder(w).Encode(data)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	data := map[string]interface{}{
		"data": "DeleteUserByID",
		"id":   vars["id"],
	}

	json.NewEncoder(w).Encode(data)
}
