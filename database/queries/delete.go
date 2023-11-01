package queries

import (
	"sf/database"
)

func Delete[M any](model *M) {
	query := database.DB.Delete(&model)
	checkErr(query.Error)
}
