package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", "test-topic", 0)
	if err != nil {
		log.Println("Error failed connectoin to kafka: ", err.Error())
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))

	b := make([]byte, 1e3)             // 10kb
	batch := conn.ReadBatch(10e3, 1e6) // 10 kb -> 1 Mb
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println("Messsages: ", string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
