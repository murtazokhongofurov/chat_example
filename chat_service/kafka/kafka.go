package kafka

import (
	"database/sql"
	"log"

	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/kafka/consumer"
)

type KafkaConn struct {
	KafkaConn *consumer.KafkaConsumer
}

type KafkaI interface {
	Reads() *consumer.KafkaConsumer
}

func NewKafkaReader(cfg config.Config, db *sql.DB) (KafkaI, func(), error) {
	kafkaReader, err := consumer.NewKafkaConsumer(cfg, db)
	if err != nil {
		log.Println("Error conn kafka: ", err.Error())
		return nil, nil, err
	}
	return &KafkaConn{
			KafkaConn: kafkaReader,
		}, func() {
			kafkaReader.Reader.Close()
		}, nil
}

func (k KafkaConn) Reads() *consumer.KafkaConsumer {
	return k.KafkaConn
}
