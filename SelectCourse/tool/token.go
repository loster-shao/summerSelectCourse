package tool

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//头
type Header struct {
	alg string `json:"alg"`
	typ string `json:"typ"`
}

func NewHeader() Header {
	return Header{
		alg: "HS256",
		typ: "JWT",
	}
}

type Payload struct {// token里面添加用户信息，验证token后可能会用到用户信息
	Iss      string `json:"iss"`      //令牌
	Exp      string `json:"exp"`      //持续时间
	IssueAt  string `json:"iat"`      //签发时间
	Username string `json:"username"` //用户名
	Uid      int                      //用户id
}

//创建token
//e30=.eyJpc3MiOiJzenMiLCJleHAiOiIxNjI4MDIwMTA1IiwiaWF0IjoiMTYyNzk4NDEwNSIsInVzZXJuYW1lIjoic3VtbWVyIiwiVWlkIjoxfQ==.iQsgFei1029s0XanBbxKdtQj9yNSz6egnhVIWG9k5Rc=
func CreateJWT(username string, id int) string {
	header := NewHeader()
	payload := Payload{
		Iss:      "szs",                                                             //
		Exp:      strconv.FormatInt(time.Now().Add(10 * time.Hour).Unix(), 10),//持续时间
		IssueAt:  strconv.FormatInt(time.Now().Unix(), 10),                    //签发时间
		Username: username,                                                          //用户名
		Uid:      id,                                                                //id
	}
	h, _ := json.Marshal(header)  //json初始化
	p, _ := json.Marshal(payload) //json初始化

	headerBase64 := base64.StdEncoding.EncodeToString(h) //[]byte->string
	payloadBase64 := base64.StdEncoding.EncodeToString(p)

	str1 := headerBase64 + "." + payloadBase64
	/*str1 := strings.Join([]string{headerBase64, payloadBase64}, ".")//字符串拼接*/

	key := "szs"

	//HMAC是密钥相关的哈希运算消息认证码，HMAC运算利用哈希算法，以一个密钥和一个消息为输入，生成一个消息摘要作为输出。
	mac := hmac.New(sha256.New, []byte(key))
	/*fmt.Println("mac", mac)*/
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	/*fmt.Println("s", s)*/
	signature := base64.StdEncoding.EncodeToString(s)

	token := /*headerBase64 + "." + payloadBase64*/ str1 + "." + signature
	return token
}

//检查token
func CheckToken(token string) (uid int, username string, err error) {
	log.Println("token:", token)
	arr := strings.Split(token, ".")//切割
	if len(arr) != 3 {
		err = errors.New("token error1")//创建err为"oken error1"
		return
	}

	_, err = base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		log.Println(err)
		err = errors.New("token error2")
		return
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error3")
		return
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		err = errors.New("token error4")
		return
	}

	str1 := arr[0] + "." + arr[1]

	key := []byte("szs")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(str1))
	s := mac.Sum(nil)
	fmt.Println("sign", sign)
	fmt.Println("ssss", s)
	if res := bytes.Compare(sign, s); res != 0 {//Compare 比较俩者是否相等
		fmt.Println("测试用代码没有卵用")
		err = errors.New("token error5")
		return
	}

	var payload Payload
	err = json.Unmarshal(pay, &payload)
	maxtime, _ := strconv.ParseInt(payload.Exp,10,64)
	if maxtime < time.Now().Unix() {
		log.Println("token验证失败")
		return
	}

	return payload.Uid, payload.Username, nil
}
