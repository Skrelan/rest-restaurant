package models

// Address is the structure for Addresses data
type Address struct {
	StreetAddress string `json:"street_address,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
	ZipCode       string `json:"zip_code,omitempty"`
}

// User is the structure for Users data
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

// Restaurant is the structure for Restraunts data
type Restaurant struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
	Address  *Address
}

// Rating is the structure for the Ratings data
type Rating struct {
	Cost               int64  `json:"cost,omitempty"`
	Food               int64  `json:"food,omitempty"`
	CleanlinessService int64  `json:"cleanliness_service,omitempty"`
	TotalScore         int64  `json:"total_score,omitempty"`
	RestaurantID       int64  `json:"restaurant_id,omitempty"`
	UserID             int64  `json:"user_id,omitempty"`
	DateCreated        string `json:"date,omitempty"`
	DateUpdated        string `json:"date,omitempty"`
}
