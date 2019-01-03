package models

import (
	"time"
)

type Test struct {
	Id int
}

type User struct {
	Id        int
	Uuid      string
	Name      string
	Account     string
	Password  string
	Online   bool
	CreatedAt time.Time
}



type Session struct {
	Id        int
	Uuid      string
	Account     string
	UserId    int
	CreatedAt time.Time
}

func (user *User) Create() (err error) {
	_, err = Db.Insert(user)
	return nil
}

func UserByAccount(account string) (user User, err error) {
	user = User{}
	_, _ = Db.Where("account=?", account).Get(&user)

	return
}
