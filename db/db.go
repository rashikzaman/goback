package db

import (
	"fmt"
	"locally/goback/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	if db == nil {
		databaseConfig := config.GetConfig().GetDatabaseConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			databaseConfig.Username,
			databaseConfig.Password,
			databaseConfig.Host,
			databaseConfig.Port,
			databaseConfig.DatabaseName)
		tempDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
