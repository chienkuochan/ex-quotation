// auth: kunlun
// date: 2019-01-24
// description:
package handle

import (
	"config"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

// 监听消息队列
func init() {
	go func() {
		messages, err := config.Channel.Consume("go.queue1", "", true, false, false, false, nil)
		if err != nil {
			fmt.Printf("Consume queue: %s faild\n", "go.queue1")
		}

		readChan := make(chan bool)

		go func() {
			for message := range messages {
				fmt.Printf("Receive message: %d - %s\n", time.Now().Unix(), message.Body)
				var data interface{}
				err = json.Unmarshal(message.Body, &data)
				if err != nil {
					fmt.Printf("Marshal mq message faild")
				}

				m := data.(map[string]interface{})
				for k, v := range m {
					if k == "cmd" {
						if v == "tick" {
							TickHandle(message.Body)
						}
					}
				}
			}
		}()
		<-readChan
	}()
}

// 推送消息
func PublishMessage(exchange, queue string, message []byte) error {
	err := config.Channel.Publish(exchange, queue, false, false, amqp.Publishing{ContentType: "text/plain", Body: message})
	if err != nil {
		return err
	}
	return nil

}
