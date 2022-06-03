package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func init() {

	// username := os.Getenv("db_user")
	// password := os.Getenv("db_pass")
	// dbName := os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")

	username := "postgres"
	password := "postgres"
	dbName := "hb_test_db"
	dbHost := "127.0.0.1"

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	err = db.Debug().AutoMigrate(
		&User{},
		&UserAuth{},
		&Event{},
		&Location{}) //Database migration
	if err != nil {
		return
	}
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
