package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	env := godotenv.Load()
	if env != nil {
		fmt.Print(env)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	port := os.Getenv("db_port")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, port, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{},&Contact{})
}

func GetDB() *gorm.DB {
	return db
}