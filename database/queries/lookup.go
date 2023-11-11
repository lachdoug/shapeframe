package queries

import (
	"sf/database"
)

func Lookup(model any) {
	database.DB.Where(model).First(model)
}
