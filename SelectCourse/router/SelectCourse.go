package router

import (
	"SelectCourse/rabbitmq"
	"SelectCourse/tool"


	_ "fmt"
	"github.com/gin-gonic/gin"

	_ "strconv"
)

func SelectCourse(c *gin.Context) {
	id,_ := c.Get("user_id")
	id_num := id.(int)
	/*id_num := tool.Get_Int_PostForm("id")*/
	course := tool.Get_Int_PostForm("course")
	rabbitmq.Order(course, id_num)
	c.JSON(200,gin.H{
		"message": "选课成功",
	})
}


