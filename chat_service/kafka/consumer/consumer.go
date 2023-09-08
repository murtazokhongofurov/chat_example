package consumer

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/storage"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader    *kafka.Reader
	ConnClose func()
	Cfg       config.Config
	Storage   storage.StorageI
}

func NewKafkaConsumer(cfg config.Config, db *sql.DB) (*KafkaConsumer, error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.KafkaHost + ":" + cfg.KafkaPort},
		Topic:     cfg.KafkaTopic,
		Partition: cfg.Partition,
		MinBytes:  1e3,
		MaxBytes:  10e6,
	})

	return &KafkaConsumer{
		Reader: reader,
		ConnClose: func() {
			reader.Close()
		},
		Cfg:     cfg,
		Storage: storage.NewStoragePg(db),
	}, nil
}

func (k *KafkaConsumer) Consume() error {
	fmt.Println("Start waiting for message >>")
	for {
		m, err := k.Reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error read message: ", err.Error())
		}
		fmt.Println("===>>> ", string(m.Value))
	}
}
