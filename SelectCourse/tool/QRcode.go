package tool

import (
	"fmt"
	"log"
	"regexp"
	"github.com/skip2/go-qrcode"
)

func CheckSQL(s string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		log.Println("error", err)
		return false
	}
	log.Println(re.MatchString(s))
	return re.MatchString(s)
}

func CheckCookie(){

}

//二维码
var png []byte
func Erweima() {
	err := qrcode.WriteFile("127.0.0.1:8080/szs", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}

func CheckPassword(password, pass string) bool {
	if password == pass{
		return true
	}
	return false
}