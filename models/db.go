package models

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB //database

func init() {
	if strings.HasSuffix(os.Args[0], ".test.exe") {
		fmt.Println("run under go test")
	} else {
		fmt.Println("normal run")
		Init()
	}
}

func Init() {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	err = db.Debug().AutoMigrate(
		&User{},
		&UserAuth{}) //Database migration
	if err != nil {
		return
	}
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	testDataBase, err := gorm.Open(sqlite.Open("./gorm_test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("db err: (TestDBInit) ", err)
	}
	sqliteDb, _ := testDataBase.DB()
	sqliteDb.SetMaxIdleConns(3)
	db = testDataBase
	return db
}

// Delete the database after running testing cases.
func TestDBFree() error {
	sqliteDb, _ := db.DB()
	err := sqliteDb.Close()
	if err != nil {
		return err
	}
	err = os.Remove("./gorm_test.db")
	return err
}
