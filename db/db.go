package db

import (
	"locally/goback/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	if db == nil {
		dsn := "host=" + config.GetConfig().GetDatabaseConfig().Host + " user=" + config.GetConfig().GetDatabaseConfig().Username +
			" password=" + config.GetConfig().GetDatabaseConfig().Password + " dbname=" + config.GetConfig().GetDatabaseConfig().DatabaseName +
			" port=" + config.GetConfig().GetDatabaseConfig().Port + " sslmode=disable TimeZone=Asia/Dhaka"
		tempDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Error(err)
			return nil
		} else {
			log.Println("Connected to database")
		}

		db = tempDb
	}
	return db
}
