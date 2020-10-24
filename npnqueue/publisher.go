package npnqueue

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/kyleu/npn/npncore"
)

type Publisher struct {
	Topic  string
	Addrs  []string
	writer sarama.SyncProducer
}

func NewPublisher(cfg *Config) (*Publisher, error) {
	config := makeSaramaConfig(cfg.Username, cfg.Password, cfg.Verbose)
	producer, err := sarama.NewSyncProducer(cfg.Addrs, config)
	if err != nil {
		panic(err)
	}

	return &Publisher{Topic: cfg.Topic, Addrs: cfg.Addrs, writer: producer}, nil
}

func (c *Publisher) Write(ctx context.Context, key string, m *Message) error {
	json := npncore.ToJSON(m.Payload, nil)
	hd := make([]sarama.RecordHeader, 0, len(m.Headers))
	for k, v := range m.Headers {
		hd = append(hd, sarama.RecordHeader{Key: []byte(k), Value: v})
	}
	message := &sarama.ProducerMessage{
		Topic:     c.Topic,
		Key:       sarama.StringEncoder(m.Key),
		Value:     sarama.StringEncoder(json),
		Headers:   hd,
		Timestamp: time.Now(),
	}
	_, _, err := c.writer.SendMessage(message)
	return err
}

func (c *Publisher) Close() error {
	return nil
}
