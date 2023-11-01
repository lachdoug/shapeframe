package queries

import (
	"fmt"
	"sf/database"
)

func Lookup(model any, assocName string, key string, value string, assoc any) {
	err := database.DB.
		Model(model).
		Where(fmt.Sprintf("%s = ?", key), value).
		Association(assocName).
		Find(assoc)
	checkErr(err)
}
