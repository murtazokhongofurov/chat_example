package kafka

import (
	"github.com/kafka_example/websocket_service/config"
	"github.com/kafka_example/websocket_service/kafka/producer"
	"github.com/kafka_example/websocket_service/pkg/logger"
)

type Kafka struct {
	KafkaFunc *producer.KafkaProducer
}

type KafkaI interface {
	Produce() *producer.KafkaProducer
}

func NewKafka(cfg config.Config, log logger.Logger) (KafkaI, func(), error) {
	kafka, err := producer.NewKafkaProducer(cfg, log)
	if err != nil {
		log.Info("Error conn kafka producer: ", logger.Error(err))
		return &Kafka{}, func() {}, err
	}
	return &Kafka{
			KafkaFunc: kafka,
		}, func() {
			kafka.ConnClose()
		}, nil

}

func (k *Kafka) Produce() *producer.KafkaProducer {
	return k.KafkaFunc
}
