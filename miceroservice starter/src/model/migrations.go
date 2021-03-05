package model

import "web-microservice-starter/utils/database"

// Migrate - Include all DB migrations that need to be run here
func Migrate() {
	db := database.GetInstance()
	db.AutoMigrate(Article{})
}
