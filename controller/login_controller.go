package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"main.go/entity"
)

func LoginController(c *gin.Context) {
	login := entity.Login{}
	er := c.ShouldBindJSON(&login)
	if er != nil {
		fmt.Println(er)
	}
	userName := Take_Username(login.Password)
	passWord := Take_Password(login.Username)
	//fmt.Println(userName, passWord, login.Username, login.Password)
	if userName == login.Username && passWord == login.Password {
		c.IndentedJSON(200, gin.H{
			"Message": "Successfully Login",
		})
	} else {
		c.IndentedJSON(200, gin.H{
			"Message": "Wrong,Plz try Again",
		})
	}
}
