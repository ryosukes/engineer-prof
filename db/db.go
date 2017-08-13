package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	toml.DecodeFile("../config/db.toml", &dbConfig)

	connection := dbConfig.user + ":" + dbConfig.pass + "@tcp([" + dbConfig.host + "]:" + dbConfig.port + ")/" + dbConfig.name + "?charset=utf8&parseTime=True"
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
