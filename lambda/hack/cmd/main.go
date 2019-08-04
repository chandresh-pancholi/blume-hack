package main

import (
	"github.com/Shopify/sarama"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"hack/pkg/lambdas"
	"log"
	"os"
	"strings"
)

func main() {

	brokers := os.Getenv("KAFKA_BROKER")
	log.Printf("brokers: %s ", brokers)
	s := session.Must(session.NewSession(aws.NewConfig()))

	s3Session := s3.New(s)

	reko := rekognition.New(s)

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy

	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	asyncProducer, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		log.Printf("async producer creation failed. %v ", err)
	}

	h := lambdas.Handler{
		S3:    s3Session,
		Rekog: reko,
		Producer: asyncProducer,
	}

	h.Successes()
	h.Errors()

	lambda.Start(h.Handle)
}
