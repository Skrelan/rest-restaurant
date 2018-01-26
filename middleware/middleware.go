package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/skrelan/LogrusWrapper/log"
	"github.com/skrelan/rest-restaurant/db"
	"github.com/skrelan/rest-restaurant/models"
)

var address []models.Address

func Health(w http.ResponseWriter, req *http.Request) {
	log.Info("Health Checkpoint pinged")
	json.NewEncoder(w).Encode("v0.01 running!")
	return
}

func AddUser(w http.ResponseWriter, req *http.Request) {

}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	limit := params.Get("limit")
	if len(limit) == 0 {
		limit = "100"
	}
	offset := params.Get("offset")
	if len(offset) == 0 {
		offset = "0"
	}

	res, err := db.GetAllUsers(&limit, &offset)
	if err != nil {
		log.Error(err)
		json.NewEncoder(w).Encode("DB connection failed")
		return
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {

}

func AddRestaurant(w http.ResponseWriter, req *http.Request) {

}

func GetRestaurants(w http.ResponseWriter, req *http.Request) {

}

func UpdateRestaurant(w http.ResponseWriter, req *http.Request) {

}

func AddReview(w http.ResponseWriter, req *http.Request) {

}

func GetReviews(w http.ResponseWriter, req *http.Request) {

}

func UpdateReviews(w http.ResponseWriter, req *http.Request) {

}

func init() {
	log.Info("Middleware started")
}
