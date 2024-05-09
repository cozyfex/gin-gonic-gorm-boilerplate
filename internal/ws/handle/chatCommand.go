package handle

import (
	"github.com/goccy/go-json"

	"gin-gonic-gorm-boilerplate/internal/repository"
	"gin-gonic-gorm-boilerplate/internal/service"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
	"gin-gonic-gorm-boilerplate/internal/ws"
)

func ChatCommand(c *ws.Client, cmd ws.Command) {
	switch cmd.Action {
	case "join":
		c.JoinRoom(cmd.Room)
	case "message":
		if c.Room != nil {
			msg := []byte("Chat: " + cmd.Message)
			c.Room.Broadcast <- msg
		}
	case "users":
		userRepo := repository.NewUserRepository(c.DB)
		userService := service.NewUserService(userRepo)

		users, err := userService.ListUser()
		if err != nil {
			return
		}

		res := make(map[string]interface{})
		res["action"] = "users"
		res["users"] = users

		usersJSON, err := json.Marshal(res)
		if err != nil {
			logger.Error(err)
		}

		c.Send <- usersJSON
	case "leave":
		c.LeaveRoom()
	}
}
