package rabbitmq

import (
	"fmt"
	"net/url"

	"github.com/YogaRP/finansial/transaction-service/internal/configs"
	"github.com/streadway/amqp"
)

const createBudgetQueue = "budget.create"

type Client struct {
	conn *amqp.Connection
}

type createBudgetMessage struct {
	UserID string `json:"user_id"`
	Limit  uint   `json:"limit"`
	Period string `json:"period"`
}

func NewClient(cfg *configs.Config) (*Client, error) {
	// rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitMQ.Username, cfg.RabbitMQ.Password, cfg.RabbitMQ.Host, cfg.RabbitMQ.Port)
	rabbitURL := url.URL{
		Scheme: "amqp",
		User:   url.UserPassword(cfg.RabbitMQ.Username, cfg.RabbitMQ.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.RabbitMQ.Host, cfg.RabbitMQ.Port),
		Path:   cfg.RabbitMQ.Vhost,
	}

	amqpUrl := rabbitURL.String()

	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

func (c *Client) DeclareQueue(queue string) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (c *Client) Consume(queue string, handler func(amqp.Delivery), prefetchCount int) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	if prefetchCount > 0 {
		if err := ch.Qos(prefetchCount, 0, false); err != nil {
			ch.Close()
			return err
		}
	}

	_, err = ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		return err
	}

	msgs, err := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		return err
	}

	go func() {
		for d := range msgs {
			handler(d)
		}
		_ = ch.Close()
	}()

	return nil
}

// func StartBudgetConsumer(client *Client, budgetService service.BudgetServiceInterface) error {
// 	if client == nil {
// 		return fmt.Errorf("rabbitmq client is nil")
// 	}

// 	return client.Consume(createBudgetQueue, func(d amqp.Delivery) {
// 		var payload createBudgetMessage
// 		if err := json.Unmarshal(d.Body, &payload); err != nil {
// 			logger.Errorf("[RabbitMQConsumer] invalid message body: %v", err)
// 			_ = d.Nack(false, false)
// 			return
// 		}

// 		userID, err := uuid.Parse(payload.UserID)
// 		if err != nil {
// 			logger.Errorf("[RabbitMQConsumer] invalid user id: %v", err)
// 			_ = d.Nack(false, false)
// 			return
// 		}

// 		request := dto.CreateBudgetRequest{
// 			UserID: userID,
// 			Limit:  payload.Limit,
// 			Period: payload.Period,
// 		}
// 		if err := request.Validate(); err != nil {
// 			logger.Errorf("[RabbitMQConsumer] invalid create budget payload: %v", err)
// 			_ = d.Nack(false, false)
// 			return
// 		}

// 		if err := budgetService.CreateBudget(context.Background(), request); err != nil {
// 			logger.Errorf("[RabbitMQConsumer] create budget failed: %v", err)
// 			_ = d.Nack(false, true)
// 			return
// 		}

// 		_ = d.Ack(false)
// 		logger.Infof("[RabbitMQConsumer] created budget for user %s", payload.UserID)
// 	}, 1)
// }
