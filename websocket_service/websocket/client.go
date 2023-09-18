package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	grpcclient "github.com/kafka_example/websocket_service/pkg/grpc_client"
)

type Client struct {
	GRPC grpcclient.GrpcClientI
}

type SocketConnection struct {
	*websocket.Conn
}

var upgradeUser = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var wsUserChan = make(chan WebscoketMessage)

var ConnectedClientList = struct {
	sync.Mutex
	connections map[SocketConnection]string
}{
	connections: make(map[SocketConnection]string),
}

func NewWebsocktConnection(c *gin.Context, grpc grpcclient.GrpcClientI) {
	ws, err := upgradeUser.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	client := &Client{GRPC: grpc}
	conn := SocketConnection{Conn: ws}
	Name := c.Query("name")

	// Connected client Insert into map
	ConnectedClientList.Mutex.Lock()
	ConnectedClientList.connections[conn] = Name
	ConnectedClientList.Mutex.Unlock()

	err = ws.WriteJSON("Welcome")
	if err != nil {
		log.Println(err)
	}
	var msg Broadcast
	msg.Type = "list_users"
	msg.Users = client.GetUserList()
	client.SendBroadcast(msg)
	go client.ListenIncoming(&conn)
	go client.HandleWebsocketMessages()
}

func (c *Client) ListenIncoming(conn *SocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", r)
		}
	}()
	var payload WebscoketMessage
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			conn.Conn.Close()
			ConnectedClientList.Lock()
			delete(ConnectedClientList.connections, *conn)
			ConnectedClientList.Unlock()
			log.Println("Client disconnected")
			break
		} else {
			payload.Conn = *conn
			wsUserChan <- payload
		}
	}
}
