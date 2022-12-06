package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/entity"
)

func Check_Special(password string) bool {
	ans := strings.ContainsAny(password, "`|~|@|#|$|%|^|&|*|(|)|_|-|+|=|>|<|,|.|?|/|'|:|;|[|]|{|}")
	return ans
}
func Check_Number(password string) bool {
	ans := strings.ContainsAny(password, "1|2|3|4|5|6|7|8|9|0")
	return ans
}
func Check_Character(password string) bool {
	ans := strings.ContainsAny(password, "A|B|C|D|E|F|G|H|I|J|K|L|M|N|O|P|Q|R|S|T|U|V|W|X|Y|Z|a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z")

	return ans
}
func Password_Check(c *gin.Context) {
	pass := entity.Passwords{}
	er := c.ShouldBindJSON(&pass)
	if er != nil {
		fmt.Println(er)
	}
	password := pass.Password
	if len(password) > 7 {
		if Check_Character(password) {
			if Check_Number(password) {
				if Check_Special(password) {
					c.IndentedJSON(200, gin.H{
						"Message": "Strong and all OK.",
					})
				} else {
					c.IndentedJSON(200, gin.H{
						"Message": "Special-Character Missing,Try Again use Special Character.",
					})
				}
			} else {
				c.IndentedJSON(200, gin.H{
					"Message": "Number Missing,Try Again use Number.",
				})
			}
		} else {
			c.IndentedJSON(200, gin.H{
				"Message": "Character Missing,Try Again use Character.",
			})
		}
	} else {
		c.IndentedJSON(200, gin.H{
			"Message": "Too Short,Try properly.",
		})
	}

}
