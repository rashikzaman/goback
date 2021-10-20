package migration

import (
	"locally/goback/db"
	"locally/goback/pkg/user/domain"
)

func Migrate() {
	db.GetDb().AutoMigrate(domain.User{})
}
