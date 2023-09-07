package producer

import (
	"context"
	"fmt"

	"github.com/kafka_example/websocket_service/config"
	"github.com/kafka_example/websocket_service/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	conn      *kafka.Conn
	ConnClose func()
}

func NewKafkaProducer(cfg config.Config, log logger.Logger) (*KafkaProducer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.KafkaHost+":"+cfg.KafkaPort, cfg.KafkaTopic, cfg.Partition)
	if err != nil {
		log.Error("Error while connection kafka: ", logger.Error(err))
	}

	return &KafkaProducer{
		conn: conn,
		ConnClose: func() {
			conn.Close()
		},
	}, nil

}

func (k *KafkaProducer) ProduceMessage(message []byte) error {
	fmt.Println("Message========>>> ", string(message))
	_, err := k.conn.WriteMessages(kafka.Message{Value: message})
	if err != nil {
		return err
	}

	fmt.Println("Successfully produce review!!!\n", string(message))
	return err
}
