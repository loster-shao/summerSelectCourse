package router

import (
	"SelectCourse/model"
	"SelectCourse/tool"
	"github.com/gin-gonic/gin"
	"log"
)

func SetCourse(c *gin.Context) {
	course := tool.GetCoursePOST()
	log.Println("course", course)
	tid, _ := c.Get("username")
	uid, _ := c.Get("user_id")
	log.Println(tid.(string), uid.(int))
	if tid.(string) == "1"{
		course.TeacherId = uid.(int)
	}
	if tid.(string) != "1" && tid.(string) != "Admin"{
		c.JSON(300, gin.H{
			"message": "权限不足",
		})
		return
	}

	if err := model.AddCourse(course); !err{
		tool.DbError_JSON()
		return
	}
	c.JSON(200,gin.H{
		"message": "添加课程成功",
	})
}

func ChangeCourse(c *gin.Context)  {

	course := tool.GetCoursePOST()
	course_id := tool.Get_Uint_PostForm("course_id")
	course.CourseId = course_id
	log.Println("courser", course_id)

	tid, _ := c.Get("username")
	uid, _ := c.Get("user_id")

	if err, _ := model.QueryUser(tid.(string), uid.(int)); err != nil{
		c.JSON(500,gin.H{
			"message": "查无此人，或课程不对",
		})
	}

	i := model.SetCourse(course);
	switch i {
	case 0:
		tool.DBSelectError_JSON()
		break
	case 1:
		tool.DbError_JSON()
		break
	case 2:
		c.JSON(200,gin.H{
			"message": "修改成功",
		})
		break
	default:
		log.Println("error")
		c.JSON(400,gin.H{
			"error": "服务器出错",
		})
		return
	}
}

func DeleteCourse(c *gin.Context) {
	tid, _ := c.Get("username")
	uid, _ := c.Get("user_id")
	course  := tool.Get_Int_PostForm("course_id")
	if err, _ := model.QueryUser(tid.(string), uid.(int)); err != nil{
		c.JSON(500,gin.H{
			"message": "权限不足",
		})
	}
	if err := model.DeleteCourse(tid.(string), uid.(int), course); !err{
		tool.DbError_JSON()
	}

	c.JSON(200, gin.H{
		"message": "success delete",
	})

}