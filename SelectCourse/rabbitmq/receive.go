package rabbitmq

import (
	"SelectCourse/model"
	"SelectCourse/struct_model"
	"SelectCourse/tool"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func OpenConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")    //连接
	tool.RabbitMQ_HandleError(err, "Can't connect to MQ")         //检测错误
	defer conn.Close()                                                 //关闭连接
	amqpChannel, err := conn.Channel()                                 //创建通道
	tool.RabbitMQ_HandleError(err, "Can't create a amqpChannel") //检测错误
	defer amqpChannel.Close()                                         //关闭通道

	//声明队列
	queue, err := amqpChannel.QueueDeclare(
		"goodList",//队列名称
		true,//持久队列
		false,//是否自动删除，没有消息就自动删除次队列
		false,//是否具有排他性（true只有自己可以访问，连接端口会自动删除）
		false,//是否阻塞（设置为true时就像没有bufio的channel）
		nil)                                                        //额外属性（我也不知道有啥额外属性）
	tool.RabbitMQ_HandleError(err, "Could not declare `add` queue") //检测错误

	err = amqpChannel.Qos(1, 0, false)                        //这个等会处理
	tool.RabbitMQ_HandleError(err, "Could not configure QoS") //检测错误

	//接受信息
	messageChannel, err := amqpChannel.Consume(
		queue.Name,//队列名称
		"",//当前消费者名称（用于区分多个消费者
		false,//是否自动应答
		false,//是否具有排他性（true只有自己可以访问，连接端口会自动删除）
		false,//[已经不支持] 如果设置为true,表示不能了个将同一个connection中发送的消息传递个ie这个connection中的消费者
		false,//消费是否阻塞
		nil,//其他参数or额外属性
	)
	tool.RabbitMQ_HandleError(err, "Could not register consumer") //检测错误
	stopChan := make(chan bool)                                   //声明nil bool channel
	//开启协程
	go func() {
		log.Println("Consumer ready, PID:", os.Getpid())//打印日志
		//遍历消息管道
		for d := range messageChannel {
			log.Println("Received a message:", string(d.Body))//打印日志
			good := &struct_model.Sc{}
			err := json.Unmarshal(d.Body, good)//序列化
			if err != nil {
				log.Println("Error decoding JSON:", err)
			}
			log.Println("course:", string(d.Body))//打印日志


			model.SelectCourse(string(d.Body))//选课


			//手动确认ACK我们可以在创建消费者的时候将auto-ack设置为false，
			// 一旦我们消费消息任务完毕的时候使用d.Ack(false)来确认ack，
			// 告诉RabbitMQ该消息可以删除
			if err := d.Ack(false); err != nil {
				log.Println("Error acknowledging message :", err)
			} else {
				log.Println("Acknowledged message")
			}
		}
	}()
	<-stopChan//发送管道
}