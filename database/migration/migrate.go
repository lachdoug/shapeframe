package migration

import (
	"sf/database"
	"sf/models"
)

func Migrate() {
	database.DB.AutoMigrate(
		&models.Shape{},
		&models.Frame{},
		&models.Workspace{},
		&models.Directory{},
		&models.UserContext{},
	)
}
