package main

import (
	"locally/goback/app/route"
	"locally/goback/db"
	"locally/goback/db/migration"
)

func main() {
	db.GetDb()
	migration.Migrate()
	route.InitRouter()
}
