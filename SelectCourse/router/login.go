package router

import (
	"SelectCourse/model"

	"SelectCourse/tool"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Login(c *gin.Context)  {

	idtype := c.PostForm("IdType")//0为学生，1为老师，2为管理员
	id_num := tool.Get_Int_PostForm("id")
	pass   := c.PostForm("password")
	log.Println(idtype, id_num, pass)

	err, ok, username := model.Login(id_num, pass, idtype)//0为学生，1为老师，2为管理员
	log.Println("errerrerr", err)
	if err != nil{
		tool.DbError_JSON()
		return
	}

	if ok == false{
		log.Println(id_num, "用户名或密码输入错误", time.Now())
		c.JSON(300, gin.H{"status": 300, "message" : "用户名or密码输入错误,请重试！"})
		return
	}
	token := tool.CreateJWT(idtype, id_num)
	c.JSON(200, gin.H{
		"status": 200,
		"message" : username + "登录成功!",
		"token": token,
	})
	/*cookie.Cookie(id_num, c)*/
}
