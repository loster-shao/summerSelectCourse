package router

import (
	"SelectCourse/model_redis"
	"time"

	"SelectCourse/tool"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"regexp"
	"strings"
)

/*var limiter = NewIPRateLimiter(1, 5)*/

//中间键测试是否是选课时间
func CheckSelectOpen(c *gin.Context)  {
	if time.Now().Unix() - tool.T.Start.Unix() > 10{
		c.Next()
	}else {
		c.JSON(404,gin.H{
			"message": "选课还未开放",
		})
		c.Abort()
	}

}

//检查token
func middle_LoginCheck(c *gin.Context)  {
	tool.C = c
	token := c.PostForm("token")
	uid, username, err := tool.CheckToken(token)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusBadGateway,gin.H{
			"message": "token验证失败",
		})
		c.Redirect(http.StatusMovedPermanently,"http://localhost:8080/login")//重定向
		c.Abort()
	}
	c.Set("username", username)
	c.Set("user_id", uid)
	c.Next()
}

func CheckSQL(c *gin.Context) {
	//log.Println("SQL")
	tool.C = c
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		log.Println("error", err)
		c.Abort()
	}

	for s, _ := range c.Request.PostForm{
		log.Println(re.MatchString(s))
		if re.MatchString(s){
			c.Abort()
			c.JSON(500,gin.H{
				"message": "数据库渗入",
			})
		}
	}
	c.Next()
}

//对恶意请求的IP进行限制
func RobotRestrict(c *gin.Context) {
	//log.Println("ip")
	//

	var (
		count = 50  // 频次数
		cycle = 60 // 时间周期 单位（second）
	)
	ip := strings.Split(c.Request.RemoteAddr, ":")[0]
	conn := model_redis.RedisPool.Get()
	/*log.Println(ip, conn)*/
	rep, _ := redis.String(conn.Do("Get", ip))
	/*log.Println(rep)*/
	if rep == "" {
		_, err := conn.Do("setex", ip, cycle, count) //conn.Do()用的是最多的，把命令行中的args一个个写进去
		if err != nil {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": 405,
					"msg":  "server error",
				},
			)
			c.Abort() //终止请求，直接返回提示信息
			return
		}
	}

	if rep == "1" {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": 405,
				"msg":  "客官慢一些",
			},
		)
		c.Abort() //终止请求，直接返回提示信息
		return
	}
	conn.Do("DECR", ip)
	defer conn.Close()

}

