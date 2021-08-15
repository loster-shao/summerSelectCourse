package tool

import (
	"SelectCourse/struct_model"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func StringToTime(times string) time.Time {
	var timestr = "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(timestr, times, time.Local)
	if err != nil {
		TypeError_JSON("time")
		return time.Time{}
	}
	return t
}

func StringToInt(str string) int {
	str_int, err := strconv.Atoi(str)
	if err != nil {
		TypeError_JSON("int")
		return -1
	}
	return str_int
}

func StringToStruct(body string) struct_model.Sc {
	var sc struct_model.Sc
	json.Unmarshal([]byte(body), &sc)
	return sc
}

func StringToUint(str string) uint {
	fmt.Println(str)
	str_uint, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		TypeError_JSON("uint")
		return 0
	}
	return uint(str_uint)
}