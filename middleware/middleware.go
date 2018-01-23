package middleware

import (
	"net/http"

	"github.com/skrelan/LogrusWrapper/log"
	"github.com/skrelan/rest-restraunt/models"
)

var address []models.Address

func AddUser(w http.ResponseWriter, req *http.Request) {

}

func GetUsers(w http.ResponseWriter, req *http.Request) {

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

func Start() {
	log.Info("Middleware started")
}
