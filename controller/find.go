package controller

import (
	"main.go/entity"
)

func Take_Password(username string) string {

	var users []entity.Register
	var p string
	Db.Table("registers").Select("password").Where("username = ?", username).Take(&users)
	for _, item := range users {
		p += string(item.Password)
	}
	return p
}

func Take_Username(password string) string {

	var users []entity.Register
	var u string
	Db.Table("registers").Select("username").Where("password = ?", password).Take(&users)
	for _, item := range users {
		u += string(item.Username)
	}
	return u
}
