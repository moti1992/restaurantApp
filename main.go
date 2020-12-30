package main

import (
	"log"
	"net/http"
	"restaurantApp/handlers"
	"restaurantApp/utils"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func init() {
	utils.ConnectToDatabase()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/search", handlers.HandleSearch).Methods(utils.ConstGET)
	log.Fatal(http.ListenAndServe(":5000", r))
	return
}
