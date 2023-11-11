package migration

import (
	"sf/database"
	"sf/models"
)

func Migrate() {
	if err := database.DB.AutoMigrate(
		&models.UserContext{},
		&models.Workspace{},
		&models.Frame{},
		&models.Shape{},
		&models.Directory{},
		&models.Configuration{},
	); err != nil {
		panic(err)
	}
}
