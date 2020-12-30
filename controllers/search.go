package controllers

import (
	"fmt"
	"net/http"
	"restaurantApp/models"

	"github.com/spf13/cast"
)

func Search(r *http.Request) (resp interface{}) {
	params := r.URL.Query()
	searchResp := models.SearchRestaurantResponseJSON{}
	result, err := models.GetRestaurantsByNameAndType(params["query"][0], cast.ToBool(params["veg"][0]))
	if err != nil {
		return searchResp
	}
	searchResp.Data.Count = len(result)
	searchResp.Data.Restaurants = result
	fmt.Println(searchResp)
	return searchResp
}
