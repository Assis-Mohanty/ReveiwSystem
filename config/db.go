package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
func SetupDb() (*sql.DB,error) {
	cfg := mysql.NewConfig()
	cfg.User=os.Getenv("DB_USER")
	cfg.Passwd=os.Getenv("DB_PASSWORD")
	cfg.Addr=os.Getenv("DB_ADDRESS")
	cfg.Net=os.Getenv("DB_NETWORK")
	cfg.DBName=os.Getenv("DB_NAME")

	// cfg.FormatDSN()
	var err error
	db,err=sql.Open("mysql",cfg.FormatDSN())
	if err != nil {
		log.Panic("Failed connecting to database")
	}
	pingErr:=db.Ping()
	if pingErr != nil {
		log.Panic("Cannot ping the database")
	}
	fmt.Println("Connected to database")

	return db,nil
}