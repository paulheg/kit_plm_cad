package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/paulheg/kit_plm_cad/backend/internal/routing"
)

type RobotID string

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})

	router := routing.NewRouter[RobotID]()

	app.Get("/robots", func(c *fiber.Ctx) error {
		return c.JSON(router.StatusList())
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/robot/:id", websocket.New(func(c *websocket.Conn) {
		id := RobotID(c.Params("id"))
		defer func() {
			log.Println("Connection of ID:", id, "was closed by the robot.")
			c.Close()
		}()

		// id cant be empty
		if len(id) == 0 {
			c.WriteMessage(websocket.CloseProtocolError, []byte("No id passed, need '/robot/:id'"))
			return
		}

		con, err := router.RegisterRobot(id, c)
		if err != nil {
			log.Println(err)
			err = c.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			if err != nil {
				log.Println(err)
			}
			return
		}

		defer func() {
			err := router.DisconnectRobot(id)
			if err != nil {
				log.Println(err)
			}
		}()

		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error while reading from robot:", err, "message:", msg)
				break
			}

			log.Printf("Read from Robot: %s", msg)

			if con.IsConnected() {
				err = con.SendToRemote(mtype, string(msg))
				if err != nil {
					log.Println("Error while sending to remote:", err)
					break
				}
			}
		}

	}))

	app.Get("/ws/remote/:id", websocket.New(func(c *websocket.Conn) {
		id := RobotID(c.Params("id"))
		defer func() {
			log.Println("Connection of ID:", id, "was closed by the remote.")
			c.Close()
		}()

		con, err := router.ConnectRemote(id, c)
		if err != nil {
			log.Println(err)
			err = c.WriteMessage(websocket.CloseMessage, []byte(err.Error()))
			if err != nil {
				log.Println(err)
			}
			return
		}

		defer func() {
			err := router.DisconnectRemote(id)
			if err != nil {
				log.Println(err)
			}
		}()

		for {
			mtype, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Error while reading from remote:", err)
				break
			}
			log.Printf("Read from Remote: %s", msg)

			if con.IsConnected() {
				err = con.SendToRobot(mtype, string(msg))
				if err != nil {
					log.Println("Error while sending to robot:", err)
					break
				}
			}

		}

	}))

	log.Fatal(app.Listen(":3000"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}
