package db

import (
	"github.com/jinzhu/gorm"
	"github.com/BurntSushi/toml"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	user string
	pass string
	host string
	name string
	port string
}

var config Config

func Connect() *gorm.DB {
	_, err := toml.DecodeFile("./../config/db.toml", &config)

	connection := config.Db.user + ":" + config.Db.pass + "@tcp([" + config.Db.host + "]:" + config.Db.port + ")/" + config.Db.name + "?charset=utf8&parseTime=True"
	db, err := gorm.Open(
		"mysql",
		connection,
	)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
