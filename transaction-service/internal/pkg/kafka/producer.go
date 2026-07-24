package kafka

import (
	"context"
	"encoding/json"

	"github.com/YogaRP/finansial/transaction-service/internal/pkg/logger"
	"github.com/segmentio/kafka-go"
)

// Producer holds the Sarama SyncProducer.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer creates a new Kafka synchronous producer.
func NewProducer(brokers []string) *Producer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
	}

	return &Producer{
		writer: writer,
	}
}

// PublishMessage sends a message to a Kafka topic.
func (p *Producer) Publish(ctx context.Context, topic, key string, value any) error {
	payload, err := json.Marshal(value)

	if err != nil {
		logger.Errorf("marshal event: %w", err)
		return err
	}

	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: payload,
	}

	if err := p.writer.WriteMessages(ctx, msg); err != nil {
		logger.Errorf("write message: %w", err)
		return err
	}

	return nil

}

// Close shuts down the producer.
func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		logger.Errorf("Failed to shut down Kafka producer cleanly: %v", err)
	}
}
