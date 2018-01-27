package main

import (
	"net/http"

	"github.com/skrelan/LogrusWrapper/log"

	"github.com/gorilla/mux"
	"github.com/skrelan/rest-restaurant/middleware"
)

func endpoints(router *mux.Router) {

	router.HandleFunc("/health", middleware.Health).Methods("GET")
	router.HandleFunc("/v1/users", middleware.GetUsers).Methods("GET")
	router.HandleFunc("/v1/users", middleware.AddUser).Methods("POST")
	router.HandleFunc("/v1/users", middleware.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/restaurants", middleware.GetRestaurants).Methods("GET")
	router.HandleFunc("/v1/restaurants", middleware.AddRestaurant).Methods("POST")
	router.HandleFunc("/v1/restaurants", middleware.UpdateRestaurant).Methods("PUT")
	router.HandleFunc("/v1/ratings", middleware.GetRatings).Methods("GET")
	router.HandleFunc("/v1/ratings", middleware.AddReview).Methods("POST")
	router.HandleFunc("/v1/ratings", middleware.UpdateReviews).Methods("PUT")
	router.HandleFunc("/v1/users/{id}", middleware.GetUsers).Methods("GET")
	router.HandleFunc("/v1/users/{id}", middleware.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/restaurants/{id}", middleware.GetRestaurants).Methods("GET")
	router.HandleFunc("/v1/restaurants/{id}", middleware.UpdateRestaurant).Methods("PUT")
	router.HandleFunc("/v1/ratings/{id}", middleware.GetRatings).Methods("GET")
	router.HandleFunc("/v1/ratings/{id}", middleware.UpdateReviews).Methods("PUT")
}

func main() {
	log.Info("Running Server on localhost:8000")
	router := mux.NewRouter()
	endpoints(router)
	http.ListenAndServe(":8000", router)
}
