package db

import (
	"fmt"

	sql "github.com/jmoiron/sqlx"

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
	db, err := sql.Open(cfg.DBType, fmt.Sprintf(params, cfg.DBName, cfg.User, cfg.Token, cfg.Host, cfg.SSLMode))
	if err != nil {
		log.Error("DB Connection failed")
		return nil, err
	}
	return db, nil
}

// GetAllUsers get's all the users
func GetAllUsers(limit, offset *string) (*[]models.User, error) {
	users := make([]models.User, 0, 0)
	db, err := cfg.startConnection()
	if db == nil {
		return nil, err
	}
	query := fmt.Sprintf(GETALLUSERS, *limit, *offset)
	log.Info("running query:", query)
	err = db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func init() {
	cfg.ready()
	log.Info("DB config set")
}
