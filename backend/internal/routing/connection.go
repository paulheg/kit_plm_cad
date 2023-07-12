package routing

import (
	"fmt"

	"github.com/gofiber/websocket/v2"
)

type Connection struct {
	remote *websocket.Conn
	robot  *websocket.Conn
}

func (c *Connection) IsConnected() bool {
	return c.remote != nil && c.robot != nil
}

func (c *Connection) SendToRemote(messageType int, message string) error {
	return c.writeMessage(messageType, message, c.remote)
}

func (c *Connection) SendToRobot(messageType int, message string) error {
	return c.writeMessage(messageType, message, c.robot)
}

func (c *Connection) writeMessage(messageType int, message string, conn *websocket.Conn) error {
	if conn != nil {
		return conn.WriteMessage(messageType, []byte(message))
	}

	return fmt.Errorf("connection not set")
}
