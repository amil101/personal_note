package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // supportive for gorm
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	searchpath := os.Getenv(("search_path"))

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s search_path=%s", dbHost, username, dbName, password, searchpath)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	db.Debug().AutoMigrate(&Note{})
}

// GetDB instance
func GetDB() *gorm.DB {
	return db
}
