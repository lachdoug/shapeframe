package migration

import (
	"sf/database"
	"sf/models"
)

func Migrate() {
	if err := database.DB.AutoMigrate(
		&models.Shape{},
		&models.Frame{},
		&models.Workspace{},
		&models.Directory{},
		&models.UserContext{},
	); err != nil {
		panic(err)
	}
}
