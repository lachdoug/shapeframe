package queries

import (
	"sf/database"
)

func Create(model any) {
	query := database.DB.Create(model)
	checkErr(query.Error)
}
