package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

func init() {
	cfg.ready()
}
