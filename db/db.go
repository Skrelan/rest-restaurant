package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	sql "github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"github.com/skrelan/LogrusWrapper/log"
	"github.com/skrelan/rest-restaurant/models"
)

var cfg configs

type configs struct {
	Host    string `json:"host"`
	DBType  string `json:"db_type"`
	DBName  string `json:"db_name"`
	User    string `json:"user"`
	Token   string `json:"token"`
	SSLMode string `json:"ssl_mode"`
}

type count struct {
	Count int `db:"count"`
}

func (cfg *configs) ready() {
	plan, err := ioutil.ReadFile("db/config.json")
	if err != nil {
		log.Error("unable to read config.json | ", err)
		return
	}
	err = json.Unmarshal(plan, cfg)
	if err != nil {
		log.Error("unable to unmarshal json | ", err)
	}
	log.Info("DB config set")
}

func (cfg *configs) startConnection() (*sql.DB, error) {
	params := "dbname=%s user=%s password=%s host=%s sslmode=%s"
	db, err := sql.Open(cfg.DBType, fmt.Sprintf(params, cfg.DBName, cfg.User, cfg.Token, cfg.Host, cfg.SSLMode))
	if err != nil {
		log.Error("DB Connection failed")
		return nil, err
	}
	return db, nil
}

// GetAllUsers gets all the users
func GetAllUsers(limit, offset *string) (*[]models.User, error) {
	users := make([]models.User, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETALLUSERS, *limit, *offset)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

// GetUserByIDs gets all the user's info
func GetUserByIDs(ids *string) (*[]models.User, error) {
	users := make([]models.User, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETUSERBYIDS, *ids)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

// InsertIntoUsers is called to create a new user
func InsertIntoUsers(user *models.User) error {
	db, err := cfg.startConnection()
	if db == nil {
		return err
	}
	query := fmt.Sprintf(INSERTINTOUSERS, user.FirstName, user.LastName, user.Phone)
	defer db.Close()
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser is called to update a user
func UpdateUser(user *models.User) error {
	var check count
	db, err := cfg.startConnection()
	if db == nil {
		return err
	}
	query := fmt.Sprintf(CHECKUSER, user.ID)
	log.Info("running query:", query)
	err = db.Get(&check, query)
	defer db.Close()
	if err != nil {
		return err
	}
	if check.Count < 1 {
		return fmt.Errorf("requested ID does not exsist")
	}
	query = fmt.Sprintf(UPDATEUSER, user.FirstName, user.LastName, user.Phone, user.ID)
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

// GetAllVenues get's all the restaurants
func GetAllVenues(limit, offset *string) (*[]models.Restaurant, error) {
	restaurants := make([]models.Restaurant, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETALLVENUES, *limit, *offset)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&restaurants, query)
	if err != nil {
		return nil, err
	}
	return &restaurants, nil
}

// InsertIntoVenues is called to create a new user
func InsertIntoVenues(restaurant *models.Restaurant) error {
	db, err := cfg.startConnection()
	if db == nil {
		return err
	}
	query := fmt.Sprintf(INSERTINTORESTAURANTS, restaurant.Name, restaurant.Category)
	defer db.Close()
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf(INSERTINTOVENUES,
		restaurant.Venue.StreetAddress,
		restaurant.Venue.City,
		restaurant.Venue.State,
		restaurant.Venue.ZipCode,
		restaurant.Name,
		restaurant.Category)
	defer db.Close()
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

// GetVenuesByIDs gets venue(s) by id(s)
func GetVenuesByIDs(ids *string) (*[]models.Restaurant, error) {
	restaurants := make([]models.Restaurant, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETVENUESBYIDS, *ids)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&restaurants, query)
	if err != nil {
		return nil, err
	}
	return &restaurants, nil
}

// GetVenuesWhere gets venue(s) by id(s)
func GetVenuesWhere(where, limit, offset *string) (*[]models.Restaurant, error) {
	restaurants := make([]models.Restaurant, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETVENUESWHERE, *where, *limit, *offset)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&restaurants, query)
	if err != nil {
		return nil, err
	}
	return &restaurants, nil
}

// UpdateVenue is called to update a Venue and/or create a new Parent Restaurant
func UpdateVenue(restaurant *models.Restaurant, updateParent bool) error {
	var check count
	var query string
	db, err := cfg.startConnection()
	if db == nil {
		return err
	}

	query = fmt.Sprintf(CHECKVENUE, restaurant.ID)
	log.Info("running query:", query)
	err = db.Get(&check, query)
	defer db.Close()
	if err != nil {
		return err
	}
	if check.Count < 1 {
		return fmt.Errorf("requested ID does not exsist")
	}
	if updateParent {
		query = fmt.Sprintf(UPDATERESTAURANT,
			restaurant.Name, restaurant.Category, restaurant.ID)
	} else {
		query = fmt.Sprintf(INSERTINTORESTAURANTS,
			restaurant.Name, restaurant.Category)
	}
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf(UPDATEVENUE,
		restaurant.Venue.StreetAddress,
		restaurant.Venue.City,
		restaurant.Venue.State,
		restaurant.Venue.ZipCode,
		restaurant.Name,
		restaurant.Category,
		restaurant.ID)
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRestaurant is called to update a parent restaurant
func UpdateRestaurant(restaurant *models.Restaurant) error {
	db, err := cfg.startConnection()
	if db == nil {
		return err
	}
	query := fmt.Sprintf(UPDATERESTAURANT, restaurant.Name, restaurant.Category, restaurant.ID)
	defer db.Close()
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

// GetAllRatings get's all the reviews by restaurant/user/venue
func GetAllRatings(limit, offset *string) (*[]models.UserRestaurantRating, error) {
	ratings := make([]models.UserRestaurantRating, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETALLRATINGS, *limit, *offset)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&ratings, query)
	if err != nil {
		return nil, err
	}
	return &ratings, nil
}

// GetRatingsWhere gets the reviews by restaurant/user/venue
func GetRatingsWhere(where, limit, offset *string) (*[]models.UserRestaurantRating, error) {
	ratings := make([]models.UserRestaurantRating, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETRATINGSWHERE, *where, *limit, *offset)
	defer db.Close()
	log.Info("running query:", query)
	err = db.Select(&ratings, query)
	if err != nil {
		return nil, err
	}
	return &ratings, nil
}

func init() {
	cfg.ready()
}

// InsertIntoRatings inserts a new rating
func InsertIntoRatings(rating *models.Rating) (string, error) {
	var check count
	db, err := cfg.startConnection()
	if db == nil {
		return "", err
	}

	query := fmt.Sprintf(CHECKUSERRATINGS,
		rating.UserID,
		rating.RestaurantID,
		time.Now().AddDate(0, 0, -30).UTC().Format("2006-01-02 15:04:05"))
	defer db.Close()
	log.Info("running query:", query)
	err = db.Get(&check, query)
	if err != nil {
		return "DB connection issue", err
	}
	log.Info(check)
	if check.Count > 0 {
		msg := "user has given a review to a venue of the restaurant in the past 30 days"
		return msg, fmt.Errorf("30 day constraint conflict")
	}
	query = fmt.Sprintf(INSERTINTORATINGS,
		rating.Cost,
		rating.Food,
		rating.CleanlinessService,
		rating.TotalScore,
		rating.RestaurantID,
		rating.UserID,
		rating.Comments,
		time.Now().UTC().Format("2006-01-02 15:04:05"), //GO Timeformat is defined as this
		time.Now().UTC().Format("2006-01-02 15:04:05"))
	defer db.Close()
	log.Info("running query:", query)
	_, err = db.Query(query)
	if err != nil {
		msg := "User has already given a rating to this venue of the restaurant"
		return msg, err
	}
	return "", nil
}
