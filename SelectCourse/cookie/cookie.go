package cookie

import "github.com/gin-gonic/gin"

func Cookie(id string, c *gin.Context)  {
	c.SetCookie(/*1*/"username", id,
		/*2*/6000,
		/*3*/"/",
		/*4*/"localhost",
		/*5*/false,
		/*6*/false)
	//第一个参数为 cookie 名；
	// 第二个参数为 cookie 值；
	// 第三个参数为 cookie 有效时长；
	// 第四个参数为 cookie 所在的目录；
	// 第五个为所在域，表示我们的 cookie 作用范围；
	// 第六个表示是否只能通过 https 访问；
	// 第七个表示 cookie 是否支持HttpOnly属性。
}
