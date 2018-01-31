package middleware

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skrelan/rest-restaurant/models"
)

var restaurants []models.Restaurant
var testData []url.Values
var solution []string

func TestRestaurantFilters(t *testing.T) {
	testData = []url.Values{
		{
			"id":   []string{"1,2,3"},
			"city": []string{"San Francisco"},
		},
		{
			"id":       []string{"1,2,3"},
			"city":     []string{"San Francisco"},
			"zip_code": []string{"94123"},
		},
		{
			"total_score": []string{"3.5"},
		},
	}
	solution = []string{
		"v.id in (1,2,3) AND LOWER(v.city) = LOWER('San Francisco')",
		"v.id in (1,2,3) AND LOWER(v.city) = LOWER('San Francisco') AND v.zip_code in ('94123')",
		"avs.score > 3.5",
	}
	assert.Equal(t, solution[0], strings.Join(*restaurantFilters(&testData[0]), " AND "))
	assert.Equal(t, solution[1], strings.Join(*restaurantFilters(&testData[1]), " AND "))
	assert.Equal(t, solution[2], strings.Join(*restaurantFilters(&testData[2]), " AND "))
}

func TestRatingFilters(t *testing.T) {
	testData = []url.Values{
		{
			"user_id": []string{"1,2,3"},
		},
		{
			"restaurant_id": []string{"1,2,3"},
		},
		{
			"user_id":       []string{"1"},
			"restaurant_id": []string{"1,2,3"},
		},
	}
	solution = []string{
		"u.id in (1,2,3)",
		"v.id in (1,2,3)",
		"u.id in (1) AND v.id in (1,2,3)",
	}
	assert.Equal(t, solution[0], strings.Join(*ratingFilters(&testData[0]), " AND "))
	assert.Equal(t, solution[1], strings.Join(*ratingFilters(&testData[1]), " AND "))
	assert.Equal(t, solution[2], strings.Join(*ratingFilters(&testData[2]), " AND "))
}
