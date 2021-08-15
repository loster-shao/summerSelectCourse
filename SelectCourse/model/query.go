package model

import (
	"SelectCourse/struct_model"

	"SelectCourse/tool"
	"errors"
	"fmt"
	"log"
)

var u struct_model.User
var t struct_model.Teacher
var a struct_model.Administrator

//查询课程
func QueryCourse(name string) ([]struct_model.Course, bool) {
	course := []struct_model.Course{}
	fmt.Println(course)
	if err := DB.Where("Course_name = ?", name).Find(&course).Error; err != nil{
		tool.DBerrorSQL(err)
		return nil, false
	}
	return course, true
}

/*func Login(id int, password, type_id string ) (error, bool, string){
	err, b, str := Query(id, password, type_id)
	fmt.Println(err, b)
	tool.CheckPassword(password, str.(struct_model.User).Password)
}*/

//登录(error有无错误， bool true继续执行， 返回一个接口或者是username
func Login(id int, password, type_id string ) (error, bool, string){

	/*id0, err := strconv.Atoi(id)
	fmt.Println(id0)
	if err != nil {
		log.Println("账号格式有误:", err)
		err := fmt.Errorf("账号格式有误")
		return err, false
	}*/
	switch type_id {
	case "0":
		if err := DB.Where("id = ?", id).First(&u).Error; err != nil{
			log.Println("u:", u)
			tool.DBerrorSQL(err)
			return err, false, ""
		}
		log.Println("u:", u)

		tool.CheckPassword(password, u.Password)
		return nil, true, u.Username
		break
	case "1":
		if err := DB.Where("teacher_id = ?", id).First(&t).Error; err != nil{
			tool.DBerrorSQL(err)
			return err, false, ""
		}
		tool.CheckPassword(password, t.Password)
		return nil, true, t.Username
		break
	case "Admin":
		if err := DB.Where("id = ?", id).First(&a).Error; err != nil{
			tool.DBerrorSQL(err)
			return err, false, ""
		}
		tool.CheckPassword(password, a.Password)
		return nil, true, a.Username
		break
	default:
		err := errors.New("格式错误")
		return err, false, ""
		break
	}
	return nil, false, ""
}

//查询用户
func QueryUser(type_id string, id int) (error, string) {
	switch type_id {
	case "0":
		if err := DB.Where("id = ?", id).First(&u).Error; err != nil{
			log.Println("u:", u)
			tool.DBerrorSQL(err)
			return err,  ""
		}
		log.Println("u:", u)

		return nil, u.Username
		break
	case "1":
		if err := DB.Where("id = ?", id).First(&t).Error; err != nil{
			tool.DBerrorSQL(err)
			return err, ""
		}

		return nil, t.Username
		break
	case "Admin":
		if err := DB.Where("id = ?", id).First(&a).Error; err != nil{
			tool.DBerrorSQL(err)
			return err, ""
		}

		return nil, a.Username
		break
	default:
		err := errors.New("格式错误")
		return err, ""
		break
	}
	return nil, ""
}

/*func Find(id int, password string, str interface{}) error {
	if err := DB.Where("user = ?", id).First(&str).Error; err != nil{
		log.Println("数据库错误", err)
		err := fmt.Errorf("数据库错误")
		return err
	}
	return nil
}*/