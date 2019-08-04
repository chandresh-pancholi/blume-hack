package api

import (
	"context"
	"log"
	"net/http"
	es "processor/pkg/elasticsearch/config"
	"processor/pkg/kafka/consumer"
)

// Server info required to run MSG as backend service
type Server struct {
	router   *http.ServeMux
	esClient *es.ESClient
}

// NewServer . create new MSG server struct
func NewServer() *Server {
	s := new(Server)
	s.router = newRouter()

	ses, err := es.NewESClient()
	if err != nil {
		log.Fatalf("eS client connection failed. error %v ", err)
	}

	s.esClient = ses

	return s
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

// Run method is to run MSG application
func (s *Server) Run() {
	s.routerHandler()

	s.initialiseConsumer()
	err := http.ListenAndServe(":8080", s.router)
	if err != nil {
		log.Fatalf("Processor starting failed. error %v", err)
		return
	}

}

func (s *Server) routerHandler() {
	s.clientHandler()
}

func (s *Server) initialiseConsumer() {
	for _, c := range []string{"blumetest" /*"food", "grocery", "electronics", "dunzo"*/} {
		log.Printf("initialising kafka consumer, topic %s, groupId %s", c, c)

		consumerGroup, handler, err := consumer.NewKafkaConsumer(c, s.esClient)
		if err != nil {
			log.Fatalf("kafka consumer group creation failed. groupId %s, error %v", c, err)
		}
		go func() {
			if err := consumerGroup.Consume(context.Background(), []string{c}, handler); err != nil {
				log.Printf("fail to consume kafka message. topic %s, error %v", c, err)
			}

		}()
	}
	log.Printf("kafka consumers up and running")

}

func (s *Server) clientHandler() {

	//s.router.HandleFunc("/client", clientHandler.CreateClient)
	//s.router.HandleFunc("/client/get", clientHandler.GetClient)
	//s.router.HandleFunc("/client/stop", clientHandler.StopClient)
}
