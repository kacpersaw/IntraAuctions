package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "intra-auctions.db")
	if err != nil {
		logrus.Error(err.Error())
		panic("Faield to connect to database")
	}

	err = db.DB().Ping()
	if err != nil {
		logrus.Panic("Ping to db failed")
	}

	logrus.Info("Connection to database established successfully!")

	return db
}
