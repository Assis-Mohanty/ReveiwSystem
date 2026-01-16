package config

import (
	"fmt"
	"log"
	"os"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func SetupDb() (*gorm.DB,error) {
	// cfg := mysql.NewConfig()
	dns:=mysqlDriver.Config{
	User:os.Getenv("DB_USER"),
	Passwd:os.Getenv("DB_PASSWORD"),
	Addr:os.Getenv("DB_ADDRESS"),
	Net:os.Getenv("DB_NETWORK"),
	DBName:os.Getenv("DB_NAME"),
	// AllowNativePasswords: true,
	Params: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
			"loc":       "Local",
		},
		
	}
	db,err:=gorm.Open(mysql.Open(dns.FormatDSN()),&gorm.Config{})
	if err != nil {
		log.Panic("Error connecting to database")
	}

	pingErr:=db.Error
	if pingErr != nil {
		log.Panic("Cannot ping the database")
	}
	fmt.Println("Connected to database")

	return db,nil
}