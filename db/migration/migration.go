package migration

import (
	"locally/goback/app/domain"
	"locally/goback/db"
)

func Migrate() {
	db.GetDb().AutoMigrate(
		domain.User{},
		domain.BusinessType{},
		domain.Business{},
	)
}
