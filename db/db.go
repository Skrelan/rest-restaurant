package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/skrelan/LogrusWrapper/log"
	"github.com/skrelan/rest-restaurant/models"
)

var cfg configs

type configs struct {
	Host    string
	DBType  string
	DBName  string
	User    string
	Token   string
	SSLMode string
}

func (cfg *configs) ready() {
	cfg.Host = "localhost"
	cfg.DBType = "postgres"
	cfg.DBName = "rest_restaurants"
	cfg.User = "postgres"
	cfg.Token = ""
	cfg.SSLMode = "disable"
}

func (cfg *configs) startConnection() (*sql.DB, error) {
	params := "dbname=%s user=%s password=%s host=%s sslmode=%s"
	db, err := sql.Open(cfg.DBType, fmt.Sprintf(params, cfg.Host, cfg.User, cfg.Token, cfg.Host, cfg.SSLMode))
	if err != nil {
		log.Error("DB Connection failed")
		return nil, err
	}
	return db, nil
}

// GetAllUsers get's all the users
func GetAllUsers() (*[]models.User, error) {
	users := make([]models.User, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}

	rows, err := db.Query(GETALLUSERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user)
		if err != nil {
			log.Warn("Record failed")
			continue
		}
		users = append(users, user)
	}
	return &users, nil
}

func init() {
	cfg.ready()
	log.Info("DB config set")
}
