package main

import (
	"locally/goback/db"
	"locally/goback/migration"
	"locally/goback/route"
)

func main() {
	db.GetDb()
	migration.Migrate()
	route.InitRouter()
}
