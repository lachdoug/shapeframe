package queries

import "sf/database"

func Save(model any) {
	query := database.DB.Save(model)
	checkErr(query.Error)
}
