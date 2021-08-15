package main

import (
	"SelectCourse/model"
	"SelectCourse/rabbitmq"
	"SelectCourse/router"
	"fmt"
	"github.com/gin-gonic/gin"

)

/*func main() {

	engines := gin.Default()
	engines.Use(RequestInfos())

	//query解析
	engines.GET("/query", func(context *gin.Context) {
		fmt.Println("中间件的使用方法")
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  context.FullPath(),
		})
	})

	engines.GET("/hello", func(context *gin.Context) {
		//
	})

	engines.Run(":8080")
}

//打印请求信息的中间件
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求Path：", path)
		fmt.Println("请求method：", method)
		fmt.Println("状态码：", context.Writer.Status())

		context.Next() //

		fmt.Println("状态码:", context.Writer.Status())
	}
}*/
func main(){
	app := gin.Default()
	model.DbConn()
	go func() {
		rabbitmq.OpenConsumer()
	}()
	fmt.Println("start")
	router.SetUpRouter(app)
	app.Run(":8080")
}

func Init()  {

}
