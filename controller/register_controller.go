package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"main.go/db"
	"main.go/entity"
)

var Db = db.ConnectDB()

func RegisterController(c *gin.Context) {
	users := entity.Register{}

	er := c.ShouldBindJSON(&users)
	if er != nil {
		fmt.Println(er)
	}
	Db.Table(users.TableName()).Create(&users)
	c.IndentedJSON(200, users)
}
