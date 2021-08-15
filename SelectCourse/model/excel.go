package model

import (
	"SelectCourse/struct_model"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
)



func ReadExcel(filename string) error {
	log.Println(filename)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")

	u := []struct_model.User{}
	for i := 0; i < len(rows); i++ {
		if i != 0{
			Register(rows[i][0], rows[i][1], rows[i][2])
			/*u = append(u, struct_model.User{
				Username: rows[i][0],
				Password: rows[i][1],
			})*/

		}

	}
	log.Println(u)


	return nil
}
