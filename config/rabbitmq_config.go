// auth: kunlun
// date: 2019-01-10
// description:
package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var Channel *amqp.Channel
var connected bool = false
var Queues []string

type Exchange struct {
	Name    string
	Type    string
	Durable bool
	Queue   []Queue
}

type Queue struct {
	Name    string
	Key     string
	Durable bool
}

type Config struct {
	Rabbitmq struct {
		Username string
		Password string
		Vhost    string
		Host     string
		Port     string
		Exchange []Exchange
	}
}

// 在类初始化之前执行，类似Java的 PostConstruct
func init() {
	mqConnection()
}

// 加载配置文件
func loadConfig() *Config {
	var config = Config{}
	configor.Load(&config, "config.json")
	return &config
}

func mqConnection() (err error) {
	config := loadConfig()
	url := "amqp://" + config.Rabbitmq.Username + ":" + config.Rabbitmq.Password + "@" + config.Rabbitmq.Host + ":" + config.Rabbitmq.Port + "/"
	if Channel == nil {
		var mqConfig = amqp.Config{ChannelMax: 10}
		conn, err = amqp.DialConfig(url, mqConfig)
		if err != nil {
			return err
		}

		Channel, err = conn.Channel()
		if err != nil {
			return err
		}

		exchange := config.Rabbitmq.Exchange
		if cap(exchange) != 0 {
			for _, e := range exchange {
				// 声明交换机
				Channel.ExchangeDeclare(e.Name, e.Type, e.Durable, false, false, true, nil)
				var queue []Queue
				queue = e.Queue
				if cap(queue) != 0 {
					for _, q := range queue {
						// 对象比较需要初始化
						if (Queue{} != q) {
							Queues = append(Queues, q.Key)
							//声明队列
							Channel.QueueDeclare(q.Name, q.Durable, false, false, false, nil)
							//绑定交换机
							Channel.QueueBind(q.Name, q.Key, e.Name, false, nil)
						}
					}
				}
			}
		}
		connected = true
		fmt.Printf("Connect rabbitmq")

	}
	return nil
}
