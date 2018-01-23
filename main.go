package main

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/skrelan/rest-restraunt/middleware"
)

func endpoints(router *mux.Router) {
	router.HandleFunc("v1/users", middleware.AddUser).Methods("POST")
	router.HandleFunc("v1/users", middleware.GetUsers).Methods("GET")
	router.HandleFunc("v1/users", middleware.UpdateUser).Methods("PUT")
	router.HandleFunc("v1/restaurants", middleware.AddRestaurant).Methods("POST")
	router.HandleFunc("v1/restaurants", middleware.GetRestaurants).Methods("GET")
	router.HandleFunc("v1/restaurants", middleware.UpdateRestaurant).Methods("PUT")
	router.HandleFunc("v1/reviews", middleware.AddReview).Methods("POST")
	router.HandleFunc("v1/reviews", middleware.GetReviews).Methods("GET")
	router.HandleFunc("v1/reviews", middleware.UpdateReviews).Methods("PUT")
}

func main() {
	middleware.Start()
	router := mux.NewRouter()
	endpoints(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
