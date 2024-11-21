package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/rabbitmq/amqp091-go"
)

type Queue interface {
	Connect() error
	PublishMessage(queue string, message interface{}) error
}

var logger = configs.GetLogger()

type rabbitMQQueue struct {
	conn *amqp091.Connection
}

func NewRabbitMQQueue() Queue {

	return &rabbitMQQueue{}
}

func (r *rabbitMQQueue) Connect() error {
	if r.conn == nil || r.conn.IsClosed() {
		queue := configs.GetEnv("QUEUE_URL")
		conn, err := amqp091.Dial(queue)
		if err != nil {
			return configs.NewError(configs.ErrInternalServer, err)
		}
		r.conn = conn
	}
	return nil
}

func (r *rabbitMQQueue) PublishMessage(queue string, message interface{}) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil, 
	)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        messageBytes,
		})
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	return nil
}
