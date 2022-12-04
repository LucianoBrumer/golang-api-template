package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data": "GetAllUsers",
	}

	json.NewEncoder(w).Encode(data)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data": "GetUserByID",
	}

	json.NewEncoder(w).Encode(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	data := map[string]interface{}{
		"data": "UpdateUserByID",
	}

	json.NewEncoder(w).Encode(data)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data": "DeleteUserByID",
	}

	json.NewEncoder(w).Encode(data)
}
