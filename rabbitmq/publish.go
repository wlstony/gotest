package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:32771/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"direct.hello", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue "+q.Name)

	err = ch.QueueBind(
		q.Name,         // queue name
		q.Name,         // routing key
		"direct.hello", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	strSlice := map[string]string{
		"test":     "aaa",
		"test1":    "bbb",
		"encoding": "I am posting gzip encoded message",
	}
	body, _ := json.Marshal(strSlice)
	err = ch.Publish(
		"direct.hello", // exchange
		q.Name,         // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			Headers: map[string]interface{}{
				"header1": "h1",
				"header2": 2,
			},
			ContentType:     "application/json",
			ContentEncoding: "gzip",
			DeliveryMode:    1,
			Priority:        9,
			CorrelationId:   "correlation-id",
			ReplyTo:         "reply-to",
			Expiration:      "10000",
			MessageId:       "message-id",
			Timestamp:       time.Now(),
			Type:            "demo",
			UserId:          "guest",//此处需要注意和连接使用的用户一致，否则消息会被拒绝
			AppId:           "app-id",
			Body:            body,
		})
	log.Printf(" [x] Sent %s", strSlice)
	failOnError(err, "Failed to publish a message")
}
