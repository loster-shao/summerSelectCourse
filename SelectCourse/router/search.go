package router

import (
	"SelectCourse/model"
	"SelectCourse/tool"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "log"

)

func QueryCourse(c *gin.Context) {
	Coursename := c.PostForm("Coursename")
	fmt.Println(Coursename)
	course, err := model.QueryCourse(Coursename)
	if !err{
		tool.DbError_JSON()
		return
	}
	c.JSON(200,gin.H{
		"message": course,
	})
}
