package middleware

import (
	"fmt"
	"net/url"
)

//LIMIT is the default limit for query results
const LIMIT string = "100"

//OFFSET is the default offset for query resuls
const OFFSET string = "0"

func restaurantFilters(params *url.Values) *[]string {
	clauses := make([]string, 0, 0)

	ids := params.Get("ids")
	if len(ids) > 0 {
		temp := fmt.Sprintf("v.id in (%s)", ids)
		clauses = append(clauses, temp)
	}
	city := params.Get("city")
	if len(city) > 0 {
		temp := fmt.Sprintf("LOWER(v.city) = LOWER('%s')", city)
		clauses = append(clauses, temp)
	}
	zipcode := params.Get("zip_code")
	if len(zipcode) > 0 {
		temp := fmt.Sprintf("v.zip_code in ('%s')", zipcode)
		clauses = append(clauses, temp)
	}
	category := params.Get("category")
	if len(category) > 0 {
		temp := fmt.Sprintf("LOWER(r.category) = LOWER('%s')", category)
		clauses = append(clauses, temp)
	}
	// totalscore := params.get("totalscore")
	// if len(ids){
	// 	temp := fmt.Sprintf("v.city in (%s)", cities)
	// 	clauses = append(clauses,temp)
	// }
	name := params.Get("name")
	if len(name) > 0 {
		temp := fmt.Sprintf("LOWER(r.name) = LOWER('%s')", name)
		clauses = append(clauses, temp)
	}
	return &clauses
}

func ratingFilters(params *url.Values) *[]string {
	clauses := make([]string, 0, 0)
	ratingID := params.Get("id")
	if len(ratingID) > 0 {
		temp := fmt.Sprintf("rate.id in (%s)", ratingID)
		clauses = append(clauses, temp)
	}
	userID := params.Get("user_id")
	if len(userID) > 0 {
		temp := fmt.Sprintf("u.id in (%s)", userID)
		clauses = append(clauses, temp)
	}
	restaurantID := params.Get("restaurant_id")
	if len(restaurantID) > 0 {
		temp := fmt.Sprintf("r.id in (%s)", restaurantID)
		clauses = append(clauses, temp)
	}
	return &clauses
}
