package database

import (
	"fmt"
	"os"

	"github.com/x64devv/assetag/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbInstanse *gorm.DB

func InitConnetion() error {

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))))
	if err != nil {
		return err
	}
	DbInstanse = db
	return nil
}

func RunMigrations() error {
	if err := DbInstanse.AutoMigrate(&models.Employee{}, &models.Asset{}, &models.Assignmet{}); err != nil {
		return err
	}
	return nil
}
