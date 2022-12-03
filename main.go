package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
)

func main() {
	r := gin.Default()

	r.POST("/register", controller.PasswordController)
	//r.POST("/login", controller.LoginController)
	r.Run(":5000")
}
