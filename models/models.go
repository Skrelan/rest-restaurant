package models

// Venue is the structure for Addresses data
type Venue struct {
	StreetAddress string `json:"street_address,omitempty" db:"street_address"`
	City          string `json:"city,omitempty" db:"city"`
	State         string `json:"state,omitempty" db:"state"`
	ZipCode       string `json:"zip_code,omitempty" db:"zip_code"`
}

// User is the structure for Users data
type User struct {
	ID        int64  `json:"id,omitempty" db:"id"`
	FirstName string `json:"first_name,omitempty" db:"first_name"`
	LastName  string `json:"last_name,omitempty" db:"last_name"`
	Phone     string `json:"phone,omitempty" db:"phone"`
}

// Restaurant is the structure for Restraunts data
type Restaurant struct {
	ID       int64  `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Category string `json:"category,omitempty" db:"category"`
	Venue    *Venue `json:"venue,omitempty" db:"a,prefix=a."`
}

// Rating is the structure for the Ratings data
type Rating struct {
	ID                 int64   `json:"id,omitempty" db:"id"`
	Cost               int64   `json:"cost,omitempty" db:"cost"`
	Food               int64   `json:"food,omitempty" db:"food"`
	CleanlinessService int64   `json:"cleanliness_service,omitempty" db:"cleanliness_service"`
	TotalScore         float64 `json:"total_score,omitempty" db:"total_score"`
	RestaurantID       int64   `json:"restaurant_id,omitempty" db:"restaurant_id"`
	UserID             int64   `json:"user_id,omitempty" db:"user_id"`
	Comments           string  `json:"comments,omitempty" db:"comments"`
	DateCreated        string  `json:"date_time_created,omitempty" db:"date_time_created"`
	DateUpdated        string  `json:"date_time_updated,omitempty" db:"date_time_updated"`
}

// UserRestaurantRating is the structure used for Ratings feed
type UserRestaurantRating struct {
	User       *User       `json:"user" db:"u,prefix=u."`
	Restaurant *Restaurant `json:"restaurant" db:"r,prefix=r."`
	Rating     *Rating     `json:"rating" db:"rate,prefix=rate."`
}
