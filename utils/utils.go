package utils

import (
	"fmt"
	"strings"

	"github.com/skrelan/rest-restaurant/models"
)

var responseCodes map[string]int
var replacer = strings.NewReplacer("'", "''") //for POSTGRES strings with ' / single quotes

func init() {
	responseCodes = map[string]int{
		"ok":                  200,
		"created":             201,
		"accepted":            202,
		"no content":          204,
		"already reported":    208,
		"bad request":         400,
		"unauth­orized":       401,
		"payment required":    402,
		"forbidden":           403,
		"not found":           404,
		"method not allowed":  405,
		"not acceptable":      406,
		"conflict":            409,
		"expect­ation failed": 417,
		"I'm a teapot":        418,
		"upgrade required":    426,
		"server issue":        500,
	}
}

type errorMessage struct {
	Error string `json:"error,omitempty"`
}

type responseMessage struct {
	Message string `json:"message,omitempty"`
}

func GenerateError(message string) *errorMessage {
	return &errorMessage{Error: message}
}

func GenerateMessage(message string) *responseMessage {
	return &responseMessage{Message: message}
}

func ResponseCodes(meaning string) int {
	return responseCodes[meaning]
}

func inRange(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func ValidateNewUser(user *models.User) error {
	if len(user.FirstName) == 0 {
		return fmt.Errorf("Invalid value for first_name")
	}
	if len(user.LastName) == 0 {
		return fmt.Errorf("Invalid value for last_name")
	}
	if len(user.Phone) != 10 {
		return fmt.Errorf("Invalid value for phone. Must be 10 digits")
	}
	return nil
}

func ValidateNewRestaurant(restaurant *models.Restaurant) error {
	if len(restaurant.Name) == 0 {
		return fmt.Errorf("Invalid value for name")
	}
	if len(restaurant.Category) == 0 {
		return fmt.Errorf("Invalid value for category")
	}
	if len(restaurant.Venue.City) == 0 {
		return fmt.Errorf("Invalid value for city")
	}
	if !inRange(len(restaurant.Venue.State), 0, 3) {
		return fmt.Errorf("Invalid value for state. Must be b/w 0 or 3 characters")
	}
	if len(restaurant.Venue.StreetAddress) == 0 {
		return fmt.Errorf("Invalid value for street_address.")
	}
	if len(restaurant.Venue.ZipCode) != 5 {
		return fmt.Errorf("Invalid value for zip_code. Must be 5 digits")
	}
	return nil
}

func ValidateNewRating(rating *models.Rating, update bool) error {
	if !inRange(int(rating.Cost), 1, 5) {
		return fmt.Errorf("Invalid value for cost. Must be int b/w 1 to 5")
	}
	if !inRange(int(rating.Food), 1, 5) {
		return fmt.Errorf("Invalid value for food. Must be int b/w 1 to 5")
	}
	if !inRange(int(rating.Cleanliness), 1, 5) {
		return fmt.Errorf("Invalid value for cleanliness. Must be int b/w 1 to 5")
	}
	if !inRange(int(rating.Service), 1, 5) {
		return fmt.Errorf("Invalid value for service. Must be int b/w 1 to 5")
	}
	if (rating.UserID <= 0) && !update {
		return fmt.Errorf("Invalid value for user_id")
	}
	if (rating.RestaurantID <= 0) && !update {
		return fmt.Errorf("Invalid value for restaurant_id")
	}
	rating.TotalScore = float64((rating.Cost + rating.Food + rating.Cleanliness + rating.Service)) / 4.0
	if (rating.TotalScore < 2) && (len(rating.Comments) == 0) {
		return fmt.Errorf("average rating is 1 and comments field is empty")
	}
	rating.Comments = replacer.Replace(rating.Comments)
	return nil
}
