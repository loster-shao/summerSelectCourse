package tool

import (
	"SelectCourse/struct_model"
	"github.com/gin-gonic/gin"
)

var C *gin.Context

//获取course PostForm
func GetCoursePOST() struct_model.Course {
	return struct_model.Course{
		TeacherId:  Get_Int_PostForm("teacher_id"),
		CourseName: C.PostForm("name"),
		Num:        Get_Int_PostForm("num"),
		Time0:      Get_Time_PostForm("time0"),
		Time1:      Get_Time_PostForm("time1"),
		Time2:      Get_Time_PostForm("time2"),
		Time3:      Get_Time_PostForm("time3"),
		Time4:      Get_Time_PostForm("time4"),
		Time5:      Get_Time_PostForm("time5"),
		Time6:      Get_Time_PostForm("time6"),
	}
}

//类型错误JSON
func TypeError_JSON(str string)  {
	C.JSON(500,gin.H{
		"error": "数据类型错误，应传输为" + str,
	})
}

//数据库JSON发送
func DbError_JSON()  {
	C.JSON(500, gin.H{
		"error": "数据库错误",
	})
}

//查询失败JSON发送
func DBSelectError_JSON()  {
	C.JSON(500,gin.H{
		"error": "未查询到此课程",
	})
}

func RabbitMQError_JSON()  {
	C.JSON(500,gin.H{
		"error": "rabbitmq err",
	})
}

/*func Select_Crouse_JSON(err string)  {
	if err != ""{
		C.JSON(500, gin.H{
			"error": err,
		})
	}
	C.JSON(200, gin.H{
		"message": "选课成功",
	})
}*/