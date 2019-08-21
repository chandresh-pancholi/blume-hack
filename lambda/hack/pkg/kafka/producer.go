package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"strings"
)

func Producer() sarama.AsyncProducer {
	brokers := os.Getenv("KAFKA_BROKER")
	log.Printf("brokers: %s ", brokers)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	asyncProducer, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		log.Printf("async producer creation failed. %v ", err)
	}

	return asyncProducer
}
