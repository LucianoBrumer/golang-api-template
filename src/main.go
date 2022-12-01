package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"golang-api/src/routes"
	"golang-api/src/utils"
)

func main() {
	utils.Config()

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	routes.InitUserRoutes(api)

	port := ":" + os.Getenv("PORT")
	fmt.Println("Server on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(port, r))
}
