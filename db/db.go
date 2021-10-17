package db

import (
	"locally/goback/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

func InitDb() {
	dsn := "host=" + config.GetConfig().GetDatabaseConfig().Host + " user=" + config.GetConfig().GetDatabaseConfig().Username +
		" password=" + config.GetConfig().GetDatabaseConfig().Password + " dbname=" + config.GetConfig().GetDatabaseConfig().DatabaseName +
		" port=" + config.GetConfig().GetDatabaseConfig().Port + " sslmode=disable TimeZone=Asia/Dhaka"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		return
	} else {
		log.Println("Connected to database")
	}
}
