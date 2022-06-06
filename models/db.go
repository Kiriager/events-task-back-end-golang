package models

import (
	"fmt"
	"time"

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
	//db.DropTableIfExists(&Contact{}, &Customer{})
	err = db.Debug().AutoMigrate(
		&User{},
		&UserAuth{},
		&Location{},
		&Event{},
	) //Database migration

	//db.Model(&Location{}).Association("Events")

	//db.Model(&Event{}).AddForeignKey("cust_id", "customers(cust_id)", "CASCADE", "CASCADE")
	//db.Model(&Event{}).AddForeignKey("cust_id", "customers(customer_id)", "CASCADE", "CASCADE") // Foreign key need to define manually

	if err != nil {
		return
	}
	//err = GetDB().Model(&Location{}).Association("Events").Error
	/*if err != nil {
		fmt.Println("__________________________________")
		fmt.Println(err)
	}*/
	//InitialDbSample()
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

func InitialDbSample() {
	locations := &[3]Location{}
	locations[0] = Location{
		Title:       "Huliapole",
		Description: "Museum",
		Latitude:    47.67791489915393,
		Longitude:   36.27254004850902,
	}
	locations[1] = Location{
		Title:       "Tokhmakh",
		Description: "Town",
		Latitude:    47.25861192153864,
		Longitude:   35.711731533173584,
	}
	locations[2] = Location{
		Title:       "Orikhiv",
		Description: "Museum",
		Latitude:    47.563391805477174,
		Longitude:   35.799495975777745,
	}

	events := &[4]Event{}
	events[0] = Event{
		Title:       "Battle",
		Description: "no description",
		Start:       time.Time{},
		End:         time.Time{},
		LocationId:  1,
	}
	events[1] = Event{
		Title:       "Fight for Ukarine",
		Description: "no description",
		Start:       time.Time{},
		End:         time.Time{},
		LocationId:  2,
	}
	events[2] = Event{
		Title:       "Dance width the Death",
		Description: "no description",
		Start:       time.Time{},
		End:         time.Time{},
		LocationId:  2,
	}
	events[3] = Event{
		Title:       "Arthilery Duel",
		Description: "no description",
		Start:       time.Time{},
		End:         time.Time{},
		LocationId:  3,
	}
	for i := 0; i < len(locations); i++ {
		err := GetDB().Create(&locations[i]).Error
		if err != nil {
			return
		}
	}
	for i := 0; i < len(events); i++ {
		err := GetDB().Create(&events[i]).Error
		if err != nil {
			return
		}
	}
}
