package models
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(101.42.149.70:3306)/queue_sys?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}), &gorm.Config{})
	if err != nil{
		return err
	}
	return err
}
