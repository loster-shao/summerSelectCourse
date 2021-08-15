package router

import (
	"SelectCourse/model"
	"SelectCourse/tool"
	"github.com/gin-gonic/gin"
)

func ChangeUser(c *gin.Context) {

	Admin_id := tool.Get_Int_PostForm("userID")
	Student_id := tool.Get_Int_PostForm("student")
	Teacher_id := tool.Get_Int_PostForm("teacher")

	model.SetUser(Admin_id, Student_id, Teacher_id)
	c.JSON(200,gin.H{
		"message": "修改成功",
	})
}

func OpenSelectCourse(c *gin.Context) {
	tool.Newtimes(tool.Get_Time_PostForm("start"), tool.Get_Time_PostForm("end"))
	/*fmt.Println(T)*/
	c.JSON(200, gin.H{
		"message": "开启成功",
	})
}
