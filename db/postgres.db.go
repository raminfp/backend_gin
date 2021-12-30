package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectPostgres() *gorm.DB {

	end := godotenv.Load(".env")
	if end != nil {
		panic("Faild to load .env file")
	}

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Tehran",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to postgres database")
	}
	//db.AutoMigrate()
	println("Connected to postgres database")

	return db
}

func ClosePostgres(db *gorm.DB) {
	dbPsql, err := db.DB()
	if err != nil {
		panic("Failed: postgres database connection")
	}
	err = dbPsql.Close()
	if err != nil {
		panic("Failed: unable to close postgre connection database")
	}
}
