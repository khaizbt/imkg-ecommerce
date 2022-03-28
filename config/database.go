package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	_ = godotenv.Load(".env")

	databaseInisialisation := "" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(databaseInisialisation), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = database

}
