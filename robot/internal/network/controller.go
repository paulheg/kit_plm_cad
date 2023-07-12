package network

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Configuration struct {
	Channel string
	Host    string
	Secure  bool
}

func (c Configuration) WebsocketURL() string {
	scheme := "ws"
	if c.Secure {
		scheme += "s"
	}

	u := url.URL{Scheme: scheme, Host: c.Host, Path: c.Channel}

	return u.String()
}

type OnMessage func(message string)

type Controller interface {
	Run(event OnMessage)
}

func New(config Configuration) Controller {
	return &controller{config: config}
}

var _ Controller = &controller{}

type controller struct {
	config Configuration
	ws     *websocket.Conn
}

func (conn *controller) Run(event OnMessage) {
	if conn.ws != nil {
		return
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		url := conn.config.WebsocketURL()

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			conn.log("connect", err, fmt.Sprintf("Cannot connect to websocket: %s", url))
			continue
		}
		conn.log("connect", nil, fmt.Sprintf("connected to websocket to %s", url))
		conn.ws = ws
		break
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for ; ; <-ticker.C {
			for {
				if conn.ws == nil {
					return
				}

				_, bytMsg, err := conn.ws.ReadMessage()
				if err != nil {
					conn.log("listen", err, "Cannot read websocket message")
					conn.Stop()
					break
				}

				message := string(bytMsg)

				conn.log("listen", nil, fmt.Sprintf("receive msg %s\n", message))
				event(message)
			}
		}
	}()

	wg.Wait()
}

func (con *controller) Stop() {
	if con.ws != nil {
		con.ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		con.ws.Close()
		con.ws = nil
	}
}

func (con *controller) log(method string, err error, message string) {
	if err != nil {
		log.Printf("error in %s: %s. err: %v", method, message, err)
	} else {
		// log.Printf("log from %s: %s", method, message)
	}

}
