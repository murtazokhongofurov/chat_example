// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pbm "github.com/kafka_example/websocket_service/genproto/message"

	grpcclient "github.com/kafka_example/websocket_service/pkg/grpc_client"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[int64]*Client

	// Inbound messages from the clients.
	broadcast chan messageChan
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	grpcclient grpcclient.GrpcClientI
}

func newHub(grpcclient grpcclient.GrpcClientI) *Hub {
	return &Hub{
		broadcast:  make(chan messageChan),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[int64]*Client),
		grpcclient: grpcclient,
	}
}

type Message struct {
	UserId  int64  `json:"user_id"`
	ChatId  int64  `json:"chat_id"`
	Message string `json:"message"`
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.userId] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.userId]; ok {
				delete(h.clients, client.userId)
				close(client.send)
			}
		case data := <-h.broadcast:
			var message Message
			err := json.Unmarshal(data.Message, &message)
			if err != nil {
				log.Println("error unmarshaling: ", err.Error())
				continue
			}
			fmt.Println("----->>> ", message)
			_, err = h.grpcclient.MessageService().AddMessage(context.Background(), &pbm.MessageReq{
				UserId:      message.UserId,
				ChatId:      message.ChatId,
				MessageText: message.Message,
			})
			if err != nil {
				log.Println(err.Error())
				continue
			}
			// cllt := h.clients[msgData.UserId]
			// select {
			// case cllt.send <- []byte(msgData.MessageText):
			// default:
			// 	close(cllt.send)
			// 	delete(h.clients, msgData.UserId)
			// }
			// for client := range h.clients {
			// 	cltt := h.clients[data.UserId]
			// 	select {
			// 	case cltt.send <- data.Message:
			// 	default:
			// 		close(cltt.send)
			// 		delete(h.clients, data.UserId)
			// 	}
			// }
		}
	}
}
