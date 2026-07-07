package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

const (
	ExchangeName = "user.exchange"
	UserCreated  = "user.created"
)

type UserCreatedPayload struct {
	UserID uuid.UUID `json:"user_id"`
}

func NewRabbitMQService(rabbitMQUrl string) (*RabbitMQService, error) {
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 1: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 2: %v", err)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		ExchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 3: %v", err)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		UserCreated,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 4: %v", err)
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		UserCreated,
		ExchangeName,
		false,
		nil,
	)

	if err != nil {
		log.Errorf("[RabbitMQService] NewRabbitMQService - 5: %v", err)
		return nil, err
	}

	return &RabbitMQService{
		conn: conn,
		ch:   ch,
	}, nil
}

func (p *RabbitMQService) PublishCreateUserBudget(ctx context.Context, routingKey string, payload UserCreatedPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("[RabbitMQService] PublishCreateUserBudget - 1: %v", err)
		return err
	}

	err = p.ch.Publish(
		ExchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)

	if err != nil {
		log.Errorf("[RabbitMQService] PublishCreateUserBudget - 2: %v", err)
		return err
	}

	return nil
}

func (p *RabbitMQService) Close() error {
	if p.ch != nil {
		p.ch.Close()
	}
	if p.conn != nil {
		return p.conn.Close()
	}
	return nil
}
