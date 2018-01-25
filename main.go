package main

import (
	"net/http"

	"github.com/skrelan/LogrusWrapper/log"

	"github.com/gorilla/mux"
	"github.com/skrelan/rest-restaurant/middleware"
)

func endpoints(router *mux.Router) {

	router.HandleFunc("/health", middleware.Health).Methods("GET")
	router.HandleFunc("/v1/users", middleware.AddUser).Methods("POST")
	router.HandleFunc("/v1/users", middleware.GetUsers).Methods("GET")
	router.HandleFunc("/v1/users", middleware.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/restaurants", middleware.AddRestaurant).Methods("POST")
	router.HandleFunc("/v1/restaurants", middleware.GetRestaurants).Methods("GET")
	router.HandleFunc("/v1/restaurants", middleware.UpdateRestaurant).Methods("PUT")
	router.HandleFunc("/v1/reviews", middleware.AddReview).Methods("POST")
	router.HandleFunc("/v1/reviews", middleware.GetReviews).Methods("GET")
	router.HandleFunc("/v1/reviews", middleware.UpdateReviews).Methods("PUT")
}

func main() {
	log.Info("Running Server on localhost:8000")
	router := mux.NewRouter()
	endpoints(router)
	http.ListenAndServe(":8000", router)
}
