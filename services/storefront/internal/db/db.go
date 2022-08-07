package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c *connection) ToString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName,
	)
}

func Init() (err error) {
	conn := connection{
		Host:     "postgres",
		Port:     "5432",
		User:     "almighty",
		Password: "secret",
		DBName:   "golang_mazaya",
	}

	DB, err = sql.Open("postgres", conn.ToString())
	if err != nil {
		return
	}

	err = DB.Ping()
	if err != nil {
		return
	}

	return nil
}
