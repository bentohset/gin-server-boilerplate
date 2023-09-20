package database

import (
	"example/gin-webserver/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Singapore", host, username, password, dbname)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to DB")
	}

	migrations(Database)
}

func SetupMockDB() {
	var err error
	dsn := "host=localhost user=postgres password=281100 dbname=testdb port=5432 sslmode=disable TimeZone=Asia/Singapore"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Successfully connected to the database")
	}

	migrations(Database)
}

func TeardownMockDB() {
	Database.Migrator().DropTable(&model.Entry{}, &model.User{}, "student_teachers")
}

func migrations(database *gorm.DB) {
	log.Println("Running Migrations")
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Entry{})

	log.Println("Migrations completed")

}
