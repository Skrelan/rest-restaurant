package middleware

import (
	"encoding/json"
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
		limit = LIMIT
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = OFFSET
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

	var whereClause string

	params := req.URL.Query()

	clauses := restaurantFilters(&params)

	limit := params.Get("limit")
	if len(limit) == 0 {
		limit = LIMIT
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = OFFSET
	}

	if len(*clauses) > 0 {
		whereClause = strings.Join(*clauses, " AND ")
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
		limit = LIMIT
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = OFFSET
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
