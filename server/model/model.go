package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kacpersaw/intra-auctions/config"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME))
	if err != nil {
		logrus.Error(err.Error())
		panic("Failed to connect to database")
	}

	err = db.DB().Ping()
	if err != nil {
		logrus.Panic("Ping to db failed")
	}

	db.AutoMigrate(&Auction{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&Bid{})

	db.LogMode(true)

	logrus.Info("Connection to database established successfully!")

	return db
}
