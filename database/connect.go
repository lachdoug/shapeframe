package database

import (
	"log"
	"os"
	"path/filepath"
	"sf/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error
	if DB, err = gorm.Open(sqlite.Open(dbFilePath()), &gorm.Config{
		Logger: fileLogger(),
	}); err != nil {
		panic(err)
	}
}

func IsExists() (is bool) {
	is = utils.IsFile(dbFilePath())
	return
}

func fileLogger() (fLog logger.Interface) {
	filePath := dbLogFilePath()
	utils.MakeFile(filePath)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	fLog = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
		LogLevel:             logger.Warn,
		ParameterizedQueries: true, // Don't include params
	})
	return
}

func dbFilePath() (dbFilePath string) {
	dbFilePath = filepath.Join(utils.DataDir("."), "sf.db")
	return
}
func dbLogFilePath() (dbLogFilePath string) {
	dbLogFilePath = filepath.Join(utils.DataDir("."), "sf.db.log")
	return
}
