package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skrelan/rest-restaurant/models"
)

func TestValidateNewUser(t *testing.T) {
	users := []models.User{
		{
			FirstName: "Tom",
			LastName:  "Jerry",
			Phone:     "4118717325",
		},
		{
			FirstName: "Tom",
			LastName:  "Jerry",
		},
		{
			FirstName: "Tom",
		},
		{},
	}
	assert.NoError(t, ValidateNewUser(&users[0]))
	assert.Error(t, ValidateNewUser(&users[1]))
	assert.Error(t, ValidateNewUser(&users[2]))
}

func TestValidateNewRestaurant(t *testing.T) {
	restaurants := []models.Restaurant{
		{
			Name:     "Sweet Tomato",
			Category: "Buffet",
			Venue: &models.Venue{
				City:          "San Mateo",
				State:         "CA",
				StreetAddress: "1234 Abc st.",
				ZipCode:       "94123"},
		},
		{
			Name:     "Sweet Tomato",
			Category: "Buffet",
			Venue: &models.Venue{
				City:          "San Mateo",
				State:         "CALIFORNIA",
				StreetAddress: "1234 Abc st.",
				ZipCode:       "94123"},
		},
		{
			Name: "Sweet Tomato",
			Venue: &models.Venue{
				City:          "San Mateo",
				State:         "CA",
				StreetAddress: "1234 Abc st.",
				ZipCode:       "94123"},
		},
		{
			Name:     "Sweet Tomato",
			Category: "Buffet",
			Venue: &models.Venue{
				StreetAddress: "1234 Abc st.",
				ZipCode:       "94123"},
		},
		{},
	}
	assert.NoError(t, ValidateNewRestaurant(&restaurants[0]))
	assert.Error(t, ValidateNewRestaurant(&restaurants[1]))
	assert.Error(t, ValidateNewRestaurant(&restaurants[2]))
	assert.Error(t, ValidateNewRestaurant(&restaurants[3]))
	assert.Error(t, ValidateNewRestaurant(&restaurants[4]))
}

func TestValidateNewRating(t *testing.T) {
	rating := []models.Rating{
		{
			Cost:         1,
			Food:         2,
			Cleanliness:  4,
			Service:      4,
			RestaurantID: 3,
			UserID:       3,
			Comments:     "A post hangover miracle!",
		},
		{
			Cost:         1,
			Food:         1,
			Cleanliness:  4,
			Service:      1,
			RestaurantID: 3,
			UserID:       3,
		},
		{},
	}
	assert.NoError(t, ValidateNewRating(&rating[0], false))
	assert.Error(t, ValidateNewRating(&rating[1], false))
	assert.Error(t, ValidateNewRating(&rating[2], false))
}
