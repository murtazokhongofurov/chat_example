package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/kafka_example/websocket_service/kafka"
)

func Run(r *gin.Engine, kafka *kafka.KafkaI) {
	hub := NewHub()
	go hub.Run()
	r.GET("/ws", func(ctx *gin.Context) {
		ServeWs(hub, ctx, kafka)
	})
}
