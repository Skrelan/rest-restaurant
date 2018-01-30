package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	json.NewEncoder(w).Encode(*utils.GenerateMessage("new user succesfully created"))
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
			w.WriteHeader(utils.ResponseCodes("bad request"))
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
	var user models.User
	var err error

	params := req.URL.Query()
	id := params.Get("id")
	if len(id) == 0 {
		id = mux.Vars(req)["id"]
	}
	if len(id) == 0 {
		// through error
		msg := "Missing user id"
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	user.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		msg := "Invalid user id"
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
	}
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		msg := "Invalid JSON passed"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
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
	err = db.UpdateUser(&user)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("conflict"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(*utils.GenerateMessage("user succesfully updated"))

}

func AddRestaurant(w http.ResponseWriter, req *http.Request) {
	var err error
	var restaurant models.Restaurant

	err = json.NewDecoder(req.Body).Decode(&restaurant)
	if err != nil {
		msg := "Invalid JSON passed"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	//check if restaurant data is good
	err = utils.ValidateNewRestaurant(&restaurant)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	err = db.InsertIntoVenues(&restaurant)
	if err != nil {
		msg := "Restaurant venue already exsists in DB"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("conflict"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	json.NewEncoder(w).Encode(*utils.GenerateMessage("new restaurant succesfully created"))

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
	var restaurant models.Restaurant
	var err error
	updateParent := false

	params := req.URL.Query()

	temp := params.Get("update_parent")
	if temp == "true" {
		updateParent = true
	}

	id := params.Get("id")
	if len(id) == 0 {
		id = mux.Vars(req)["id"]
	}
	if len(id) == 0 {
		// through error
		msg := "Missing restaurant venue id"
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	restaurant.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		msg := "Invalid restaurant id"
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
	}
	err = json.NewDecoder(req.Body).Decode(&restaurant)
	if err != nil {
		msg := "Invalid JSON passed"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	//check if user is good
	err = utils.ValidateNewRestaurant(&restaurant)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	err = db.UpdateVenue(&restaurant, updateParent)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("conflict"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(*utils.GenerateMessage("restaurant succesfully updated"))

}

func AddRating(w http.ResponseWriter, req *http.Request) {

	var err error
	var rating models.Rating

	err = json.NewDecoder(req.Body).Decode(&rating)
	if err != nil {
		msg := "Invalid JSON passed"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	log.Info(rating)
	//check if restaurant data is good
	err = utils.ValidateNewRating(&rating, false)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	msg, err := db.InsertIntoRatings(&rating)
	if err != nil {
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("conflict"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	json.NewEncoder(w).Encode(*utils.GenerateMessage("new restaurant succesfully created"))

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

func UpdateRating(w http.ResponseWriter, req *http.Request) {
	var rating models.Rating
	var err error

	params := req.URL.Query()

	id := params.Get("id")
	if len(id) == 0 {
		id = mux.Vars(req)["id"]
	}
	if len(id) == 0 {
		// through error
		msg := "Missing restaurant venue id"
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	rating.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		msg := "Invalid restaurant id"
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
	}
	err = json.NewDecoder(req.Body).Decode(&rating)
	if err != nil {
		msg := "Invalid JSON passed"
		log.Error(msg, err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(msg))
		return
	}
	//check if user is good
	err = utils.ValidateNewRating(&rating, true)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("bad request"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	err = db.UpdateRating(&rating)
	if err != nil {
		log.Error(err)
		w.WriteHeader(utils.ResponseCodes("conflict"))
		json.NewEncoder(w).Encode(*utils.GenerateError(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(*utils.GenerateMessage("rating succesfully updated"))

}

func init() {
	log.Info("Middleware started")
}
