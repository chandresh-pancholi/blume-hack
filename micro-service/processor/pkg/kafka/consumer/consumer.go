package consumer

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"log"
	"processor/model"
	elasticsearch "processor/pkg/elasticsearch/config"
	"processor/workflow/detect"
)

type GroupHandler struct {
	esClient *elasticsearch.ESClient
}

// NewKafkaConsumer creates a new kafka consumer
func NewKafkaConsumer(groupID string, esClient *elasticsearch.ESClient) (sarama.ConsumerGroup, *GroupHandler, error) {
	kafkaConfig := sarama.NewConfig()

	kafkaConfig.Version = sarama.V2_1_0_0
	kafkaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.ClientID = groupID //Client Id
	kafkaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, groupID, kafkaConfig)
	if err != nil {
		return nil, nil, err
	}

	h := GroupHandler{
		esClient: esClient,
	}

	return consumerGroup, &h, nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (c *GroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var detectTextOutput model.DetectText
		err := json.Unmarshal(message.Value, &detectTextOutput)
		if err != nil {
			log.Fatalf("json unmarshalling failed. error %v ", err)
		}
		//log.Printf("kafka message claimed", zap.String("value", string(message.Value)), zap.Int64("offset", message.Offset), zap.String("key", string(message.Key)))

		dw := detect.DetectWorkflow{
			EsClient: c.esClient,
		}
		log.Printf("triggering detect workflow ")
		dw.Trigger(detectTextOutput)

		sess.MarkMessage(message, "")
	}
	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *GroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	log.Printf("setup kafka consumer group. memerID %v, GenerationID %v Claims %v ", s.MemberID(), s.GenerationID(), s.Claims())
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *GroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	log.Printf("cleanup consumer group", zap.String("MemberID", s.MemberID()), zap.Int32("GenerationID", s.GenerationID()), zap.Any("Claims", s.Claims()))
	return nil
}
