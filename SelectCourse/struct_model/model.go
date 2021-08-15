package struct_model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Course struct{
	CourseId    uint       `gorm:"primary_key; autoTncrement"`
	TeacherId   int       `json:"teacher_id"`
	CourseName  string    `json:"course_name"`
	Num         int       `json:"num"`
	Time0       time.Time `json:"time_0"`
	Time1       time.Time `json:"time_1"`
	Time2       time.Time `json:"time_2"`
	Time3       time.Time `json:"time_3"`
	Time4       time.Time `json:"time_4"`
	Time5       time.Time `json:"time_5"`
	Time6       time.Time `json:"time_6"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Sex      bool
}

type Sc struct {
	Id       int `gorm:"primaryKey;autoIncrement:false"`
	CourseId int `gorm:"primaryKey;autoIncrement:false"`
}

type Teacher struct {
	TeacherId int `gorm:"primary_key"`
	Username  string
	Password  string
	Sex       bool
}

type Administrator struct {
	Id      int    `gorm:"primary_key"`
	Username string
	Password string
}

type Times struct{
	Start time.Time
	End   time.Time
}

