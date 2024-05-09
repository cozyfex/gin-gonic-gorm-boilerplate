package ws

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"

	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
)

type Client struct {
	WS            *websocket.Conn
	Send          chan []byte
	Room          *Room
	HandleCommand func(*Client, Command)
	DB            *database.Manager
}

func NewClient(ws *websocket.Conn, db *database.Manager, room *Room, handleCommand func(*Client, Command)) *Client {
	return &Client{
		WS:            ws,
		Send:          make(chan []byte, 256),
		Room:          room,
		DB:            db,
		HandleCommand: handleCommand,
	}
}

// readPump listens for new messages from the websocket connection
func (c *Client) readPump() {
	defer func() {
		c.Room.Unregister <- c
		err := c.WS.Close()
		if err != nil {
			logger.Error(fmt.Sprintf("error: %v", err))
			return
		}

		if c.Room.IsEmpty() {
			hub := GetHubInstance()
			logger.Info(fmt.Sprintf("Room %s is empty, deleting...\n", c.Room.Name))
			delete(hub.Rooms, c.Room.Name)
		}
	}()

	for {
		_, message, err := c.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error(fmt.Sprintf("error: %v", err))
			}
			break
		}

		var cmd Command
		if err := json.Unmarshal(message, &cmd); err != nil {
			logger.Error(fmt.Sprintln("Error unmarshalling message:", err))
			continue
		}

		c.HandleCommand(c, cmd)
	}
}

// writePump sends messages from the send channel to the websocket connection
func (c *Client) writePump() {
	for {
		message, ok := <-c.Send
		if !ok {
			err := c.WS.WriteMessage(websocket.CloseMessage, []byte{})
			if err != nil {
				logger.Error(fmt.Sprintf("error: %v", err))
				return
			}
			return
		}

		err := c.WS.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			logger.Error(fmt.Sprintf("error: %v", err))
			return
		}
	}
}

// JoinRoom handles joining a new room
func (c *Client) JoinRoom(roomName string) {
	if c.Room != nil {
		c.Room.Unregister <- c
	}

	hub := GetHubInstance()
	if _, exists := hub.Rooms[roomName]; !exists {
		hub.Rooms[roomName] = NewRoom(roomName)
	}

	c.Room = hub.Rooms[roomName]
	c.Room.Register <- c
}

// LeaveRoom handles leaving the current room
func (c *Client) LeaveRoom() {
	if c.Room != nil {
		c.Room.Unregister <- c
		c.Room = nil
	}
}
