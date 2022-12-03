package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/entity"
)

func Check_Special(p string) bool {
	res := strings.ContainsAny(p, "`|~|!|@|#|$|%|^|&|*|(|)|_|-|+|=|>|<|,|.|?|/|'|;|:|]|[")

	return res
}

func Check_Number(p string) bool {
	res := strings.ContainsAny(p, "1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 0")

	return res
}

func Check_Character(p string) bool {
	res1 := strings.ContainsAny(p, "A | B | C | D | E | F | G | H | I |J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z")
	res2 := strings.ContainsAny(p, "a | b | c | d | e | f | g | h | i | j | k | l | m | n | o | p | q | r | s | t | u | v | w | x | y | z")

	if res1 || res2 {
		return true
	}
	return false
}
func PasswordController(c *gin.Context) {
	pass := entity.PasswordCheck{}

	er := c.ShouldBindJSON(&pass)
	if er != nil {
		fmt.Println(er)
	}
	p := pass.Password

	if len(string(p)) > 6 {
		if Check_Character(p) == true {
			if Check_Number(p) == true {
				if Check_Special(p) == true {
					c.IndentedJSON(200, gin.H{
						"Message": "Strong and All Ok!",
					})
				} else {
					c.IndentedJSON(200, gin.H{
						"Message": "Try again and Include special character.",
					})
				}
			} else {
				c.IndentedJSON(200, gin.H{
					"Message": "Try again and Include Number.",
				})
			}
		} else {
			c.IndentedJSON(200, gin.H{
				"Message": "Try again and Include Character.",
			})
		}
	} else {
		c.IndentedJSON(200, gin.H{
			"Message": "Must be lenght 6+",
		})
	}

}
