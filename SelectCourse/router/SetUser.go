package router

import (
	"SelectCourse/model"
	"SelectCourse/tool"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context)  {
	id := tool.Get_Int_PostForm("id")
	idtype := tool.Get_Int_PostForm("idtype")

	str, i, err := model.DeleteUser(id, idtype);
	if err != nil{
		tool.DBerrorSQL(err)
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
		str:  i,
	})
}
