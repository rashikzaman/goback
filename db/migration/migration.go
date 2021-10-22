package migration

import (
	"locally/goback/db"
	"locally/goback/pkg/domain"
)

func Migrate() {
	db.GetDb().AutoMigrate(
		domain.User{},
		domain.BusinessType{},
		domain.Business{},
	)
}
