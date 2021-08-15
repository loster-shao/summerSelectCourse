package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(sapp *gin.Engine)  {
	//公共接口
	sapp.Use(CheckSQL, RobotRestrict)//OK

	//注册
	sapp.POST("/register", Register)//OK
	sapp.POST("/login", Login)//OK


	app := sapp.Group("/test")
	app.Use(middle_LoginCheck)//OK
	//管理员与教师公共接口
	app.POST("/QueryCourse", QueryCourse)//OK查询课程
	app.POST("/SetCourse", SetCourse)//OK创建课程
	app.POST("/ChangeCourse", ChangeCourse)//OK修改课程
	app.POST("/DeleteCourse", DeleteCourse)//OK删除课程
	sapp.POST("/ExcelRegister", ExcelRegister)//OK 通过上传excel进行创建用户
	//管理员接口
	app.POST("/ChangeUser", ChangeUser)
	//TODO 不知道这个有什么能修改的。。。所以下面写了个删除user 这接口没啥用

	app.POST("DeleteUser", DeleteUser)//OK删除用户
	app.POST("/OpenSelectCourse", OpenSelectCourse)//TODO 记得测试
	//学生接口
	app.POST("/SelectCourse", CheckSelectOpen, SelectCourse)//TODO 记得测试

	//教师接口
	//TODO excel表格录入（明天写完）已完成表格上传，表格以可以正常读取，逻辑部分2h左右 OK
	//TODO 测试报告 2h内应该能完成 （不知道这个怎么写啊emm）
	//TODO ip请求速率限制（明天写完）已经写好代码1h左右测试   OK
	//TODO 重定向：当登录失效或未登录时可以重定向到登录界面 ：网页行不行没有测试 OK
	fmt.Println("")
}



//公共接口
//login
//register

//管理员接口
//SetCourse
//ChangeCourse
//DeleteCourse
//ChangeUser
//OpenCourse

//学生接口
//SelectCourse
//QueryCourse

//教师接口
//SetCourse
//ChangeCourse
//DeleteCourse
