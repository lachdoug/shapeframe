package queries

import "sf/database"

func Update(model any) {
	query := database.DB.Updates(model)
	checkErr(query.Error)
}
