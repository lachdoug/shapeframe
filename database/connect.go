package database

import (
	"log"
	"os"
	"sf/app/dirs"
	"sf/app/errors"
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
	var config logger.Config
	if errors.Debug {
		config = logger.Config{
			LogLevel:             logger.Info,
			ParameterizedQueries: false,
		}
	} else {
		config = logger.Config{
			LogLevel:             logger.Error,
			ParameterizedQueries: true,
		}
	}
	fLog = logger.New(log.New(file, "\r\n", log.LstdFlags), config)
	return
}

func dbFilePath() (dbFilePath string) {
	dbFilePath = dirs.WorkspaceDir("sf.db")
	return
}
func dbLogFilePath() (dbLogFilePath string) {
	dbLogFilePath = dirs.WorkspaceDir("sf.db.log")
	return
}
