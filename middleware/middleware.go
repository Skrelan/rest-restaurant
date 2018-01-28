package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/skrelan/LogrusWrapper/log"
	"github.com/skrelan/rest-restaurant/db"
	"github.com/skrelan/rest-restaurant/models"
	"github.com/skrelan/rest-restaurant/utils"
)

func Health(w http.ResponseWriter, req *http.Request) {
	log.Info("Health Checkpoint pinged")
	json.NewEncoder(w).Encode("v0.01 running!")
	return
}

func AddUser(w http.ResponseWriter, req *http.Request) {
	var err error
	var user models.User

	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		msg := "Invalid JSON passed"
		log.Error(msg, err)
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	//check if user is good
	err = utils.ValidateNewUser(&user)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	err = db.InsertIntoUsers(&user)
	if err != nil {
		msg := "User already exsists in DB"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("conflict"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
}

func GetUsers(w http.ResponseWriter, req *http.Request) {
	var res *[]models.User
	var err error

	params := req.URL.Query()
	ids := params.Get("ids")
	if len(ids) == 0 {
		ids = mux.Vars(req)["id"]
	}
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

	memory := make([]string, 0, 0)
	clauses := &memory // because GOlang doesn't allocate memory to an empty slice
	var whereClause string

	params := req.URL.Query()

	id := mux.Vars(req)["id"]
	if len(id) > 0 {
		temp := fmt.Sprintf("r.id = %s", id)
		*clauses = append(*clauses, temp)
	} else {
		clauses = restaurantFilters(&params)
	}

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
	var res *[]models.UserRestaurantRating
	var err error
	var whereClause string

	memory := make([]string, 0, 0)
	clauses := &memory // because GOlang doesn't allocate memory to an empty slice

	params := req.URL.Query()

	id := mux.Vars(req)["id"]
	if len(id) > 0 {
		temp := fmt.Sprintf("rate.id = %s", id)
		*clauses = append(*clauses, temp)
	} else {
		clauses = ratingFilters(&params)
	}

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
		res, err = db.GetRatingsWhere(&whereClause, &limit, &offset)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode("DB connection failed")
			return
		}
	} else {
		res, err = db.GetAllRatings(&limit, &offset)
		if err != nil {
			log.Error(err)
			json.NewEncoder(w).Encode("DB connection failed")
			return
		}
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateReviews(w http.ResponseWriter, req *http.Request) {

}

func init() {
	log.Info("Middleware started")
}
