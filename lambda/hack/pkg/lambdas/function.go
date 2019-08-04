package lambdas

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	reko2 "hack/pkg/reko"
	"log"
)

type Handler struct {
	S3    *s3.S3
	Rekog *rekognition.Rekognition
	Producer sarama.AsyncProducer
}

func (h Handler) Handle(ctx context.Context, e events.S3Event) ([]rekognition.DetectTextOutput, error) {
	result := make([]rekognition.DetectTextOutput, 0)

	reko := reko2.Reko{
		Rekog: h.Rekog,
	}
	for _, record := range e.Records {
		bucket := record.S3.Bucket.Name

		key := record.S3.Object.Key

		detectOutput, err := reko.DetectText(bucket, key)
		if err != nil {
			log.Printf("detect text processing failed. bucket %s, key %s ", bucket, key)
		}
		if detectOutput != nil {
			result = append(result, *detectOutput)
		}
		do, err := json.Marshal(detectOutput)
		if err != nil {
			log.Fatalf("object to json conversion failed")
		}
		h.Publish(string(do), bucket, key)
	}

	return result, nil
}

//Publish is to emit event to Kafka
func (h Handler) Publish(payload, topic, key string) {
	h.Producer.Input() <- &sarama.ProducerMessage{
		Topic:    topic,
		Key:      sarama.StringEncoder(key),
		Value:    sarama.StringEncoder(payload),
	}
}

//Errors keep  the track of failed messages.
func (h *Handler) Errors() {
	go func() {
		for err := range h.Producer.Errors() {
			_, errEncode := err.Msg.Key.Encode()
			if errEncode != nil {
				log.Printf("error %v ", errEncode)
			}
		}
	}()
}

//Successes is to check if message successfully delivered to kafka
func (h *Handler) Successes() {
	go func() {
		for msg := range h.Producer.Successes() {
			_, errEncode := msg.Key.Encode()
			if errEncode != nil {
				log.Printf("error %v ", errEncode)
			}
		}
	}()
}
