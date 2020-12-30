package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/spf13/cast"
)

type Restaurant struct {
	RestaurantID int     `orm:"column(restaurant_id);auto" json:"restaurant_id"`
	Name         string  `orm:"column(name)" json:"name"`
	URL          string  `orm:"column(url)" json:"url"`
	Cuisines     string  `orm:"column(cuisines)" json:"cuisines"`
	Image        string  `orm:"column(image)" json:"image"`
	Address      string  `orm:"column(address)" json:"address"`
	City         string  `orm:"column(city)" json:"city"`
	Rating       float64 `orm:"column(rating)" json:"rating"`
	Veg          bool    `orm:"column(veg)" json:"veg"`
}

type SearchRestaurantResponseJSON struct {
	Data struct {
		Count       int          `json:"count"`
		Restaurants []Restaurant `json:"restaurants"`
	} `json:"data"`
}

func (t *Restaurant) Restaurant() string {
	return "restaurant"
}

func init() {
	orm.RegisterModel(new(Restaurant))
}

func GetRestaurantsByNameAndType(query string, isVeg bool) (v []Restaurant, err error) {
	o := orm.NewOrm()
	var r []orm.Params
	_, err = o.Raw(`SELECT 
			T0.restaurant_id,
			T0.name,
			T0.url,
			T0.cuisines,
			T0.image,
			T0.address,
			T0.city,
			T0.rating,
			T0.veg
		FROM
			restaurant T0
		WHERE
			(T0.name LIKE '%`+query+`%' OR T0.cuisines LIKE '%`+query+`%' OR T0.city LIKE '%`+query+`%') and veg=?
		ORDER BY T0.restaurant_id ASC`, isVeg).Values(&r)

	if err != nil {
		return nil, err
	}

	for index, _ := range r {
		var restaurant Restaurant
		restaurant.RestaurantID = cast.ToInt(r[index]["restaurant_id"])
		restaurant.Name = cast.ToString(r[index]["name"])
		restaurant.URL = cast.ToString(r[index]["url"])
		restaurant.Cuisines = cast.ToString(r[index]["cuisines"])
		restaurant.Image = cast.ToString(r[index]["image"])
		restaurant.Address = cast.ToString(r[index]["address"])
		restaurant.City = cast.ToString(r[index]["city"])
		restaurant.Rating = cast.ToFloat64(r[index]["rating"])
		restaurant.Veg = cast.ToBool(r[index]["veg"])
		v = append(v, restaurant)
	}

	return v, nil
}
