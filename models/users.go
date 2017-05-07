package userModel

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ryosukes/engineer-prof/db"
)

type User struct {
	Id          int64
	Name        string
	Email       string
	RoleId      int
	Img         string
	IsGravatar  int
	AccessToken string
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

var user User
var users []User

func (u *User) Migrate() {
	db := db.Connect()
	db.AutoMigrate(&User{})
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) All(limit int, offset int) []User {
	db := db.Connect()
	db.AutoMigrate(&User{})
	db.Limit(limit).Offset(offset).Order("id desc,name").Find(&users)
	return users
}

func (u *User) Fetch(id int64) User {
	db := db.Connect()
	user = User{Id: id}
	db.First(&user)
	return user
}

func (u *User) MapByName(name string) []User {
	db := db.Connect()
	db.Where(map[string]interface{}{"name": name}).Find(&users)
	return users
}
