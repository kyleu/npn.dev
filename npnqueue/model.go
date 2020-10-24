package npnqueue

import (
	"github.com/Shopify/sarama"
	"github.com/kyleu/npn/npncore"
	"log"
	"os"
	"strings"
	"time"
)

type Message struct {
	Topic   string            `json:"topic"`
	Key     string            `json:"key"`
	Headers map[string][]byte `json:"headers,omitempty"`
	Payload string            `json:"payload"`
	Time    time.Time         `json:"time,omitempty"`
}

type Messages []*Message

type Config struct {
	Addrs    []string
	Username string
	Password string
	Topic    string
	Verbose  bool
}

func LoadConfig(verbose bool) *Config {
	addrs := strings.Split(env("host", "pkc-4nym6.us-east-1.aws.confluent.cloud:9092"), ",")
	u := env("username", "KSLPPOL5ACWKDXJK")
	p := env("password", "TxMxQY091Tz53E7t4VBPl/nI4ZFCVUzvViX/S4dLKEcu2JkcIuGdM/zE5mPtaXhB")
	topic := env("topic", "Nuevo")
	return &Config{Addrs: addrs, Username: u, Password: p, Topic: topic, Verbose: verbose}
}

func env(key string, dflt string) string {
	x, ok := os.LookupEnv(npncore.AppKey + "_kafka_" + key)
	if ok && len(x) > 0 {
		return x
	}
	return dflt
}

func makeSaramaConfig(username string, password string, verbose bool) *sarama.Config {
	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	config := sarama.NewConfig()
	config.ClientID = "nuevo"
	config.Net.SASL.Enable = true
	config.Net.SASL.User = username
	config.Net.SASL.Password = password
	config.Net.TLS.Enable = true
	config.Producer.Return.Successes = true
	return config
}

