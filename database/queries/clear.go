package queries

import (
	"sf/database"
)

func Clear(model any, assocName string) {
	err := database.DB.Model(model).Association(assocName).Clear()
	checkErr(err)
}
