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
		"reset content":       205,
		"partial content":     206,
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

func GenerateError(message string) *errorMessage {
	return &errorMessage{Error: message}
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
