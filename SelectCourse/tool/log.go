package tool

import "log"

func DBerrorSQL(err error)  {
	log.Println("数据库错误", err)
}

func RabbitMQ_HandleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		RabbitMQError_JSON()
	}
}
