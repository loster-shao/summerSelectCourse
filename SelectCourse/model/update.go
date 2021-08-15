package model

import (
	"SelectCourse/struct_model"

	"SelectCourse/tool"
	"log"
)

func SetCourse(course struct_model.Course) int {
	/*if err := DB.Where("course_id = ?", course.CourseId).Error;err != nil{
		log.Println("select_error:", err)
		testss.DBerrorSQL(err)
		return 0
	}*/
	if err := DB.Model(&course).Where("course_id = ?", course.CourseId).Updates(&course).Error; err != nil{
		log.Println("save_error:", err)
		tool.DBerrorSQL(err)
		return 1
	}
	return 2
}

func SetUser(...int)  {

}







/*//创建房间
func Create(play1, play2, num string) error {
	fmt.Println("7")
	DB.AutoMigrate(&Room{})
	err := DB.CreateTable(&Room{
		Play1: play1,
		Play2: play2,
		Num:   num,
	}).Error
	fmt.Println("8")
	if err != nil{
		fmt.Println("err", err)
	}
	return nil
}

//加入房间
func Join(user, room string) error {
	r := &Room{
		Play2: user,
		Num:   room,
	}
	err := DB.Model(&r).Update("play2", room).Error
	return err
}*/
