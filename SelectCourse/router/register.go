package router

import (
	"SelectCourse/model"
	"SelectCourse/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func Register(c *gin.Context)  {
	idtype := c.PostForm("IdType")//0学生，1老师，2管理员
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == ""{
		fmt.Println("账号名或密码不能为空")
		c.JSON(300, gin.H{"status": 300, "message" : "用户名或密码不能为空"})
		return
	}
	num := model.Register(username, password, idtype);
	switch num {
	case 0:
		c.JSON(500,gin.H{
			"message": "用户已存在",
		})
		break

	case 1:
		tool.DbError_JSON()
		break

	case 2:
		c.JSON(200,gin.H{
			"message": "注册成功",
		})
	}
}

func ExcelRegister(c *gin.Context)  {
	f, err := c.FormFile("f1")  //根据name返回给第一个文件
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}else {
		//将读取到的文件保存到本地(服务端)
		//dst := fmt.Sprintf("./%s", f.Filename)
		dst := path.Join("./", f.Filename)//拼接字符串作为文件的路径
		_    = c.SaveUploadedFile(f, dst)//核心代码，将拿到的文件存储到指定位置
		model.ReadExcel(dst)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}

}


