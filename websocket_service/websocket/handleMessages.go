package websocket

import (
	"fmt"

	"github.com/gin-gonic/gin"

	grpcclient "github.com/kafka_example/websocket_service/pkg/grpc_client"
)

func Run(r *gin.Engine, grpc grpcclient.GrpcClientI) {
	r.GET("/ws", func(c *gin.Context) {
		NewWebsocktConnection(c, grpc)
	})
}

func (c *Client) HandleWebsocketMessages() {
	for {
		select {
		case e := <-wsUserChan:
			fmt.Println(e)
			switch e.Type {
			case "left":
				delete(ConnectedClientList.connections, e.Conn)
				var response Response
				response.Type = "list_users"
				response.Users = c.GetUserList()
				c.SendBroadcast(response)
			case "broadcast":
				go func(message WebscoketMessage) {
					var response Response
					response.Type = "broadcast"
					user := c.GetName(e)
					response.Message = fmt.Sprintf("<strong>%s</strong>: %s", user, e.Message)
					c.SendBroadcast(response)
				}(e)

			}
		}

	}
}

func (c *Client) SendInduvidual(name string, data interface{}) error {
	ConnectedClientList.Mutex.Lock()
	defer ConnectedClientList.Unlock()

	for conn, storedName := range ConnectedClientList.connections {
		if storedName == name {
			err := conn.WriteJSON(data)
			if err != nil {
				fmt.Println("Error writing JSON: ", err)
				delete(ConnectedClientList.connections, conn)
				return err
			}
		}
	}
	return nil
}

func (c *Client) SendBroadcast(data interface{}) {
	ConnectedClientList.Mutex.Lock()
	defer ConnectedClientList.Unlock()

	for conn := range ConnectedClientList.connections {

		err := conn.WriteJSON(data)
		if err != nil {
			fmt.Println("Error writing JSON: ", err)
			delete(ConnectedClientList.connections, conn)
		}
	}
}

func (c *Client) GetUserList() []string {
	var userlist []string
	for _, x := range ConnectedClientList.connections {
		if x != "" {
			userlist = append(userlist, x)
		}
	}
	return userlist
}

func (c *Client) GetName(e WebscoketMessage) string {
	var user string
	for conn, name := range ConnectedClientList.connections {
		if conn == e.Conn {
			user = name
		}
	}
	return user
}
