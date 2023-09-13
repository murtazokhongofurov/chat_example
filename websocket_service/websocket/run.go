package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/kafka_example/websocket_service/kafka"
	grpcclient "github.com/kafka_example/websocket_service/pkg/grpc_client"
)

func Run(r *gin.Engine, kafka *kafka.KafkaI, client grpcclient.GrpcClientI) {
	hub := newHub(client)
	go hub.Run()
	r.GET("/ws", func(ctx *gin.Context) {
		ServeWs(hub, ctx, kafka, client)
	})
}
