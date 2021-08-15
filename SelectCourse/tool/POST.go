package tool

import (
	"fmt"
	"time"
)

func Get_Uint_PostForm(str string) uint {
	s := C.PostForm(str)
	fmt.Println("daf", str)
	s_uint := StringToUint(s)
	return s_uint
}

func Get_Int_PostForm(str string) int {
	s := C.PostForm(str)
	str_int := StringToInt(s)
	return str_int
}

func Get_Time_PostForm(str string) (time.Time) {
	t := C.PostForm(str)
	if t == "" {
		return time.Unix(0,0)
	}
	str_time := StringToTime(t);
	return str_time
}
