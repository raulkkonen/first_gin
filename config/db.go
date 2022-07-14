package config

import "gorm.io/gorm"
import "gorm.io/driver/postgres"
import "github.com/raulkkonen/first_gin/models"

var DB *gorm.DB
func Connect(){
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}