package queries

import (
	"sf/database"
)

func Append(model any, assocName string, append any) {
	err := database.DB.Model(model).Association(assocName).Append(append)
	checkErr(err)
}
