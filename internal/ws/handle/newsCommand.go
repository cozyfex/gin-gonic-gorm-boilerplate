package handle

import "gin-gonic-gorm-boilerplate/internal/ws"

func NewsCommand(c *ws.Client, cmd ws.Command) {
	switch cmd.Action {
	case "join":
		c.JoinRoom(cmd.Room)
	case "message":
		if c.Room != nil {
			msg := []byte("News: " + cmd.Message)
			c.Room.Broadcast <- msg
		}
	case "leave":
		c.LeaveRoom()
	}
}
