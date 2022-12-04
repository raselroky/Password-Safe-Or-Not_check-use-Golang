package controller

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/entity"
)

func Check(s, s1 string) bool {
	p := false
	ans := strings.Split(s, ".")
	for _, item := range ans {
		ans1 := strings.Split(item, " ")
		for _, item1 := range ans1 {
			if item1 == s1 {
				p = true
				break
			}
		}
	}
	return p
}
func Search_Sentence(c *gin.Context) {
	sentence := entity.Search{}

	er := c.ShouldBindJSON(&sentence)
	if er != nil {
		fmt.Println(er)
	}
	s := sentence.Sentence
	s1 := sentence.Searchs
	p1 := strings.Contains(s, s1)
	if p1 {
		c.IndentedJSON(200, gin.H{
			"Message": "Yes,  This word stable in this Sentence.",
		})
	} else if p1 == false {
		if Check(s,s1) {
			c.IndentedJSON(200, gin.H{
				"Message": "Yes,  This word stable in this Sentence.",
			})
		} else {
			c.IndentedJSON(200, gin.H{
				"Message": "No,  This word not stable in this Sentence.",
			})
		}
	} else {
		c.IndentedJSON(200, gin.H{
			"Message": "No,  This word not stable in this Sentence.",
		})
	}

}
