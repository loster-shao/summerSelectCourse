package model

import (
	"SelectCourse/struct_model"
	"SelectCourse/tool"
	"github.com/jinzhu/gorm"
	"log"
	_ "time"
)

//注册
func Register(username, password, type_id string) int {
	log.Println(username, password, type_id)
	var u struct_model.User
	var t struct_model.Teacher
	var a struct_model.Administrator

	/*id0, err := strconv.Atoi(id)
	fmt.Println(id0)
	if err != nil {
		log.Println("账号格式有误:", err)
		err := fmt.Errorf("账号格式有误")
		return err, false
	}*/
	switch type_id {
	case "0":
		err := DB.Select("username").Where("username = ?", username).First(&u).Error;
		log.Println("user:", u)
		if err != nil{
			log.Println("查询用户：", err)
		}
		if u != (struct_model.User{}) {
			log.Println("用户存在")
			return 0
		}
		DB.AutoMigrate(&struct_model.User{})
		if err := DB.Create(&struct_model.User{
			Model:    gorm.Model{},
			Username: username,
			Password: password,
			Sex:      false,
		}).Error; err != nil{
			return 1
		}
		return 2

	case "1":
		err := DB.Table("teacher").Select("username").Where("username = ?", username).First(&t).Error;
		log.Println("user:", t)
		if err != nil{
			log.Println("查询用户：", err)

		}
		if t != (struct_model.Teacher{}) {
			log.Println("用户存在,t:", t)
			return 0
		}
		DB.AutoMigrate(&struct_model.Teacher{})
		if err := DB.Create(&struct_model.Teacher{
			Username: username,
			Password: password,
			Sex:      false,
		}).Error; err != nil{
			log.Println("教师无法创建")
			return 1
		}
		return 2

	case "Admin":
		err := DB.Table("Admin").Select("username").Where("username = ?", username).First(&a).Error;
		log.Println("user:", a)
		if err != nil{
			log.Println("查询用户：", err)
		}
		if a != (struct_model.Administrator{}) {
			log.Println("用户存在")
			return 0
		}
		DB.AutoMigrate(&struct_model.Administrator{})
		if err := DB.Create(&struct_model.Administrator{
			Username: username,
			Password: password,

		}).Error; err != nil{
			return 1
		}
		return 2
	default:
		return 1
	}
	return 1
}

//添加课程
func AddCourse(course struct_model.Course) bool {
	var c  struct_model.Course
	err := DB.Select("course_id").Where("course_id = ?", course.CourseId).First(&c).Error;
	if err != nil{
		log.Println("err：", err)
	}
	if c != (struct_model.Course{}) {
		log.Println("课程存在", course)
		return false
	}
	log.Println(course)
	DB.AutoMigrate(&struct_model.Course{})
	if err := DB.Table("course").Create(&course).Error; err != nil{
		tool.DBerrorSQL(err)
		return false
	}
	return true
}

func SelectCourse(str string)  {
	log.Println("select", str)
	sc := tool.StringToStruct(str)
	i := 0
	var course struct_model.Course

	/*err := DB.Select("course_id").Where("course_id = ?", sc.CourseId).Find(&sc).Error*/
	err := DB.Select("num").Where("course_id = ?", sc.CourseId).Find(&course).Error
	log.Println(course.Num, i)
	if err != nil{
		tool.DBerrorSQL(err)

		return
	}
	if i < course.Num {
		DB.AutoMigrate(&struct_model.Sc{})
		if err := DB.Create(struct_model.Sc{
			Id:       sc.Id,
			CourseId: sc.CourseId,
		}).Error; err != nil{
			tool.DBerrorSQL(err)

			return
		}
	}else {

		return
	}
	return
}
