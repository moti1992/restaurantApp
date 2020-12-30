package handlers

import (
	"encoding/json"
	"net/http"
	"restaurantApp/controllers"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	resp := controllers.Search(r)
	json.NewEncoder(w).Encode(resp)
}
