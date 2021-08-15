package tool

import (

	"SelectCourse/struct_model"
	"time"
)
var T struct_model.Times

func Newtimes(start, end time.Time) {
	T = struct_model.Times{
		Start: start,
		End:   end,
	}
}

/*func (t *times)Clock() {
	for i:=0; i<0 ; i++ {
		if t.start.Unix() - time.Now().Unix() < 10{
			True = true
			t.UnClock()
			break;
		}else {

		}
		time.Sleep(5000)
	}
}*/

/*func (t *times)UnClock()  {
	for i:=0; i<0 ; i++ {
		if time.Now().Unix() - t.end.Unix() <= 0{
			True = false
			t.Clock()
			break;
		}
		time.Sleep(5000)
	}
}*/



