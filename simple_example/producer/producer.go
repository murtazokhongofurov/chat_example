package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", "test-topic", 0)
	if err != nil {
		log.Println("Error failed connectoin to kafka: ", err.Error())
	}
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	_, err = conn.WriteMessages(
		kafka.Message{
			Value: []byte("Hello New Kafka!!")},
		kafka.Message{
			Value: []byte("Salom yangi kafka!!")},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
