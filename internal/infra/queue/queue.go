package queue

type Queue interface {
	PublishMessage(queue string, message string) error
}

type rabbitMQQueue struct {
}

func NewRabbitMQQueue() Queue {
	return &rabbitMQQueue{}
}

func (r *rabbitMQQueue) PublishMessage(queue string, message string) error {
	return nil
}
