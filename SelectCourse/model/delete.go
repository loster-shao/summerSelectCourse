package model

import (
	"SelectCourse/struct_model"

	"SelectCourse/tool"
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm"
)

func DeleteCourse(str string, userID, courseID int) bool {
	var course struct_model.Course
	fmt.Println(courseID)

	if userID != 0 {
		err := DB.Select("course_id").Where("course_id = ?", courseID).Find(&course).Error;
		if err != nil{
			log.Println("Select_error:", err)
			tool.DBerrorSQL(err)
			return false
		}

	}

	if err := DB.Where("course_id = ?", courseID).Delete(&course).Error; err != nil{
		tool.DBerrorSQL(err)
		return false
	}
	return true
}

func DeleteUser(id, idtype int) (string, int, error) {
	var u struct_model.User
	var t struct_model.Teacher
	log.Println(id, idtype)

	if idtype == 0{
		err := DB.Table("user").Select("id").Where("id = ?", id).Find(&u).Error;
		if err != nil{
			tool.DBerrorSQL(err)
			tool.DbError_JSON()
			return "", 0, err
		}
		if err := DB.Table("user").Where("id = ?", id).Delete(&u).Error; err != nil{
			tool.DBerrorSQL(err)
			tool.DbError_JSON()
			return "", 0, err
		}
		return "学生", int(u.ID), nil
	}

	if idtype == 1 {
		err := DB.Table("teacher").Select("teacher_id").Where("teacher_id = ?", id).Find(&t).Error;
		if err != nil{
			tool.DBerrorSQL(err)
			tool.DbError_JSON()
			return "", 0, err
		}
		if err := DB.Table("teacher").Where("teacher_id = ?", id).Delete(&t).Error; err != nil{
			tool.DBerrorSQL(err)
			tool.DbError_JSON()
			return "", 0, err
		}
		return "老师", t.TeacherId, err
	}

	return "", 0, nil
}


/*func DeleteGoodsCart(username string, id int) bool {
	fmt.Println(username, id)
	if err := DB.Where("id=?, username=?",id,username).Delete(&Goods{}).Error; err != nil{
		testss.DBerrorSQL(err)
		return false
	}
	return true
}*/
