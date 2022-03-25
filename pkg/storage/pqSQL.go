package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Table Const
const (
	ItemsTable  = "\"Items\""
	OrdersTable = "\"Orders\""
)

//Columns Const
const (
	orderIdColumns = "order_id"
)

type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewPgsqlDB(cfg ConfigDB) (*sqlx.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	db, err := sqlx.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
