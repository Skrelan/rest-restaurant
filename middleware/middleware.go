package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/skrelan/LogrusWrapper/log"
	"github.com/skrelan/rest-restaurant/db"
	"github.com/skrelan/rest-restaurant/models"
)

func Health(w http.ResponseWriter, req *http.Request) {
	log.Info("Health Checkpoint pinged")
	json.NewEncoder(w).Encode("v0.01 running!")
	return
}

func AddUser(w http.ResponseWriter, req *http.Request) {

}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	var res *[]models.User
	var err error

	params := req.URL.Query()
	ids := params.Get("ids")

	limit := params.Get("limit")
	if len(limit) == 0 {
		limit = "100"
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = "0"
	}
	if len(ids) > 0 {
		res, err = db.GetUserByIDs(&ids)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode("DB connection failed")
			return
		}
	} else {
		res, err = db.GetAllUsers(&limit, &offset)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode("DB connection failed")
			return
		}
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {

}

func AddRestaurant(w http.ResponseWriter, req *http.Request) {

}

func GetRestaurants(w http.ResponseWriter, req *http.Request) {
	var res *[]models.Restaurant
	var err error
	clauses := make([]string, 0, 0)
	var whereClause string

	params := req.URL.Query()
	ids := params.Get("ids")
	if len(ids) > 0 {
		temp := fmt.Sprintf("v.id in (%s)", ids)
		clauses = append(clauses, temp)
	}
	city := params.Get("city")
	if len(city) > 0 {
		temp := fmt.Sprintf("LOWER(v.city) in LOWER('%s')", city)
		clauses = append(clauses, temp)
	}
	zipcode := params.Get("zip_code")
	if len(zipcode) > 0 {
		temp := fmt.Sprintf("v.zip_code in ('%s')", zipcode)
		clauses = append(clauses, temp)
	}
	category := params.Get("category")
	if len(category) > 0 {
		temp := fmt.Sprintf("LOWER(r.category) in LOWER('%s')", category)
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

	limit := params.Get("limit")
	if len(limit) == 0 {
		limit = "100"
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = "0"
	}

	if len(clauses) > 0 {
		whereClause = strings.Join(clauses, " AND ")
		res, err = db.GetVenuesWhere(&whereClause, &limit, &offset)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode("DB connection failed")
			return
		}
		json.NewEncoder(w).Encode(res)
	} else {
		res, err = db.GetAllVenues(&limit, &offset)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode("DB connection failed")
			return
		}
		json.NewEncoder(w).Encode(res)
	}
}
func UpdateRestaurant(w http.ResponseWriter, req *http.Request) {

}

func AddReview(w http.ResponseWriter, req *http.Request) {

}

func GetRatings(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	limit := params.Get("limit")
	if len(limit) == 0 {
		limit = "100"
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = "0"
	}

	res, err := db.GetAllRatings(&limit, &offset)
	if err != nil {
		log.Error(err)
		json.NewEncoder(w).Encode("DB connection failed")
		return
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateReviews(w http.ResponseWriter, req *http.Request) {

}

func init() {
	log.Info("Middleware started")
}
