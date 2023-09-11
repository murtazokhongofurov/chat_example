package consumer

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "sync" 

    "github.com/kafka_example/chat_service/config"
    "github.com/kafka_example/chat_service/pkg/logger"
    "github.com/kafka_example/chat_service/storage"
    "github.com/kafka_example/chat_service/storage/models"
    "github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
    conn      *kafka.Conn
    strg      storage.StorageI
    ConnClose func()
    wg        sync.WaitGroup 
}

func NewKafkaConsumer(cfg config.Config, log logger.Logger, strg storage.StorageI) (*KafkaConsumer, error) {
    conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.KafkaHost+":"+cfg.KafkaPort, cfg.KafkaTopic, 0)
    if err != nil {
        log.Error("Error while connecting to Kafka: ", logger.Error(err))
        return nil, err 
    }

  
    wg := sync.WaitGroup{}

    return &KafkaConsumer{
        conn: conn,
        strg: strg,
        ConnClose: func() {
            conn.Close()
        },
        wg: wg, 
    }, nil
}

func (c *KafkaConsumer) ConsumeMessages() {
    defer c.wg.Done() 

    for {
        msg, err := c.conn.ReadMessage(10e3)
        if err != nil {
            log.Println("Error while reading message from Kafka: ", logger.Error(err))
            continue
        }


        c.wg.Add(1) 
        go func(message kafka.Message) {
            defer c.wg.Done() 

            var person models.Message
            if err := json.Unmarshal(message.Value, &person); err != nil {
                log.Println("Error unmarshaling message: ", logger.Error(err))
                return
            }

            fmt.Println(person)

           
            if err := c.strg.ChatApp().SavePerson(person); err != nil {
                log.Println("Error saving person: ", logger.Error(err))
            }
        }(msg)
    }
}

func (c *KafkaConsumer) Close() {
    c.conn.Close()
    c.wg.Wait() 
}
