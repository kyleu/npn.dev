package npnqueue

import (
	"context"

	"github.com/Shopify/sarama"
	"logur.dev/logur"
)

type Consumer struct {
	Topic   string `json:"topic,omitempty"`
	Addrs   []string
	Reader  sarama.ConsumerGroup
	Group   string
	Handler *ConsumeHelper
	Config  *sarama.Config
	logger  logur.Logger
}

func NewConsumer(cfg *Config, group string, handler *ConsumeHelper) (*Consumer, error) {
	config := makeSaramaConfig(cfg.Username, cfg.Password, cfg.Verbose)
	r, err := sarama.NewConsumerGroup(cfg.Addrs, group, config)
	if err != nil {
		return nil, err
	}
	return &Consumer{Topic: cfg.Topic, Addrs: cfg.Addrs, Reader: r, Group: group, Config: config, Handler: handler}, nil
}

func (c *Consumer) ReadLoop(onMessage func(msg *Message), onError func(e error)) {
	go func() {
		for {
			err := c.Read(context.Background())
			if err != nil {
				panic(err)
			}
			c.Reset()
		}
	}()

	for {
		select {
		case msg := <-c.Handler.Consumers:
			go onMessage(msg)
		case err := <-c.Handler.Errors:
			go onError(err)
		}
	}
}

func (c *Consumer) Read(ctx context.Context) error {
	return c.consume(ctx)
}

func (c *Consumer) Close() error {
	// return c.Reader.Close()
	return nil
}

func (c *Consumer) Reset() {
	c.Handler.Ready = make(chan bool)
}

func (c *Consumer) consume(ctx context.Context) error {
	return c.Reader.Consume(ctx, []string{c.Topic}, c.Handler)
}

type ConsumeHelper struct {
	Consumers chan *Message
	Errors    chan error
	Ready     chan bool
}

func NewConsumeHelper() *ConsumeHelper {
	return &ConsumeHelper{Consumers: make(chan *Message), Errors: make(chan error), Ready: make(chan bool)}
}

func (c *ConsumeHelper) Setup(sarama.ConsumerGroupSession) error {
	close(c.Ready)
	return nil
}

func (c *ConsumeHelper) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ConsumeHelper) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for m := range claim.Messages() {
		hd := make(map[string][]byte, len(m.Headers))
		for _, v := range m.Headers {
			hd[string(v.Key)] = v.Value
		}
		msg := &Message{
			Topic:   m.Topic,
			Key:     string(m.Key),
			Headers: hd,
			Payload: string(m.Value),
			Time:    m.Timestamp,
		}
		c.Consumers <- msg
		session.MarkMessage(m, "")
	}

	return nil
}
