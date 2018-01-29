package utils

import (
	"fmt"

	"github.com/skrelan/rest-restaurant/models"
)

var responseCodes map[string]int

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
	if (len(restaurant.Venue.State) == 0) || (len(restaurant.Venue.State) > 3) {
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
