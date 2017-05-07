package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryosukes/engineer-prof/config"
)

func Connect() *gorm.DB {
	connection := dbConfig.DbUser + ":" + dbConfig.DbPass + "@tcp([" + dbConfig.DbHost + "]:" + dbConfig.DbPort + ")/" + dbConfig.DbName + "?charset=utf8&parseTime=True&loc=" + dbConfig.DbLocate
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
