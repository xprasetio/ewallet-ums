package helpers

import (
	"ewallet-ums/internal/models"
	"fmt"
	"log"

	logrus "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupMySQL(){
	
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		GetEnv("DB_USER", "root"), GetEnv("DB_PASS", ""), GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_NAME", "ewallet_ums"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	logrus.Info("database connected")
	DB.AutoMigrate(&models.User{}, &models.UserSession{})

}