package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/pkg/logger"
	"github.com/kafka_example/chat_service/storage"
	"github.com/kafka_example/chat_service/storage/models"
	"github.com/segmentio/kafka-go"
)



type KafkaConsumer struct {
	conn  *kafka.Conn
	strg storage.StorageI
	ConnClose  func()
	messageChan chan kafka.Message
}

func NewKafkaConsumer(cfg config.Config, log logger.Logger, strg storage.StorageI )(*KafkaConsumer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.KafkaHost+":"+cfg.KafkaPort, cfg.KafkaTopic, 0 )
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &KafkaConsumer{
		conn: conn,
		strg: strg,
		messageChan: make(chan kafka.Message),

	}, nil
}

func(c *KafkaConsumer) ConsumeMessages(){
	go c.handleMessage()
	for {
		msg, err := c.conn.ReadMessage(10e3)
		if err != nil {
			log.Println("Error reading Kafka message", err.Error())
			continue
		}
		c.messageChan <- msg
	}
}

func (c *KafkaConsumer) handleMessage(){
	for msg := range c.messageChan {
		fmt.Println("Message received: ", string(msg.Value))
		var message models.Message
		err := json.Unmarshal(msg.Value, &message)
		if err != nil {
			fmt.Println("Error unmarshalling message", err.Error())
		}
		err = c.strg.ChatApp().SavePerson(message)
		if err != nil {
			fmt.Println("Error saving message", err.Error())
		}else {
			fmt.Println("Successfully saved message")
		}
	}
}

