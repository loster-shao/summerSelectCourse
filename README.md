##选课
http://localhost:8080/test/SelectCourse

当发送该POST请求后向rabbitmq队列中发送选课信息。

Rabbitmq服务端使用协程进行接受请求发来的消息，并进行处理（将数据添加到选课表中）
并向前端发送json

##防DDoS攻击

使用redis数据库，对ip进行注册，使其在规定时间内只能操作一定次数，每操作一次便会使得
该ip在剩余时间内的操作次数-1，到0为止，如果为0则直接返回，避免抢课期间请求过多导致软件
崩溃。

##防SQL注入

使用了正则对于SQL常用的insert，update等数据进行过滤，如发现则中断该ip的请求。


## 注册   
http://localhost:8080/register  
对于多个角色（老师学生与管理员的分辨采用id_type这一参数并使用了switch进行分别处理。
下文中的通过上传excel注册与登录也是如此
```json
{
			"name": "注册",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "IdType",
							"value": "0",
							"type": "text"
						},
						{
							"key": "username",
							"value": "test",
							"type": "text"
						},
						{
							"key": "password",
							"value": "123456",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:8080/register",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"register"
					]
				}
			},
			"response": []
		}
```

##上传excel文件进行注册
```json
{
			"name": "上传注册",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "formdata",
					"formdata": [
						{
							"key": "f1",
							"type": "file",
							"src": "/C:/Users/24565/Desktop/test.xlsx"
						}
					]
				},
				"url": {
					"raw": "http://localhost:8080/ExcelRegister",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"ExcelRegister"
					]
				}
			},
			"response": []
		}
```

## 登录
http://localhost:8080/login

登录的同时创建token并传输到前端进行保存，每次在执行需要登录的操作时，后端接受
前端传来的token进行解析。

```json
		{
			"name": "登录",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "urlencoded",
					"urlencoded": [
						{
							"key": "IdType",
							"value": "0",
							"type": "text"
						},
						{
							"key": "id",
							"value": "1",
							"type": "text"
						},
						{
							"key": "password",
							"value": "123456",
							"type": "text"
						}
					]
				},
				"url": {
					"raw": "http://localhost:8080/login",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"login"
					]
				}
			},
			"response": [ 
                                               {
                         				"code": 200,
                         				"name": "OK"
                         			}
                             ]
		}

```   

##token

创建token：token中包含了id_type（用户的类别，上文有提到）与用户id，以防止随意被人篡改，
进行攻击等。

在解析token的同时会c.SET()以上两值，以供后续函数进行读取操作，防止了随意篡改的可能。 

##二维码

采用了网上所给库，因为没有加入手机扫码的服务，所以并未正式使用。

##tool包

json的发送（我不知道json发送放入其中是否合适）：
在检查token的同时创建了一个全局变量C（拷贝于 gin.Context)
这样可以减少代码量，比如数据库的错误可以多次使用该返回json函数

为了方便赋值，对于string与int等等的转化问题上，为了减少golang中常常出现的判断err值的问题，
使用了如下的方法：
```
func StringToInt(str string) int {
	str_int, err := strconv.Atoi(str)
	if err != nil {
		TypeError_JSON("int")
		return -1
	}
	return str_int
}
```
以此简化了代码。

同样，对于x-www-form-urlencoded所传输的string，经常会需要转int等各种格式，
所以我写了如下函数进行直接赋值：
```
func Get_Int_PostForm(str string) int {
	s := C.PostForm(str)
	str_int := StringToInt(s)//该函数参考上一函数
	return str_int
}
```
如此便可直接给予赋值，简化代码

##

由于json接口有点多，也有点长，如果想看可以移步阅读 

***测试.postman_collection.json***

接口文档

尾声：这段时间参加三下乡，时间匆忙，代码中存在的bug与疏漏肯定有许多，
希望不要介意，感谢阅读。



