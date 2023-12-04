package server

import (
	"context"
	"encoding/json"
	"log"
	"tgbot-service/internal/biz"
	"tgbot-service/internal/conf"
	"tgbot-service/pkg/rabbit"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageDTO struct {
	Message string `json:"message"`
}

func NewRabbitConn(c *conf.Data, uc *biz.UserUseCase) *rabbit.RabbitConn {
	conn, err := amqp.Dial(c.Rabbit)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	ch.QueueDeclare(
		"social", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	msgs, err := ch.Consume(
		"social",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Не удалось зарегистрировать consumer: %s", err)
	}
	go func() {
		for d := range msgs {
			var message MessageDTO
			err := json.Unmarshal(d.Body, &message)
			if err == nil {
				uc.SendAll(context.TODO(), message.Message)
			}
		}
	}()

	return rabbit.NewRabbitConn(conn, ch)
}
