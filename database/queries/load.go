package queries

import (
	"sf/database"

	"gorm.io/gorm"
)

func Load(model any, id uint, preloads ...string) {
	db := database.DB
	for _, pl := range preloads {
		db = db.Preload(pl)
	}
	query := db.First(model, id)
	if query.Error == gorm.ErrRecordNotFound {
		return
	}
	checkErr(query.Error)
}
