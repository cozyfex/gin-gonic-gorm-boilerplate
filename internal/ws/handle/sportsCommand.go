package handle

import "gin-gonic-gorm-boilerplate/internal/ws"

func SportsCommand(c *ws.Client, cmd ws.Command) {
	switch cmd.Action {
	case "join":
		c.JoinRoom(cmd.Room)
	case "message":
		if c.Room != nil {
			msg := []byte("Sports: " + cmd.Message)
			c.Room.Broadcast <- msg
		}
	case "leave":
		c.LeaveRoom()
	}
}
