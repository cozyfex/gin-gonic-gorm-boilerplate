package routing

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/ws"
)

func WithDB(db *database.Manager, h func(c *gin.Context, db *database.Manager)) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c, db)
	}
}

func WithWebSocket(db *database.Manager, hub *ws.Hub, roomName string, handleCommand func(c *ws.Client, cmd ws.Command)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ws.HandleConnections(c.Writer, c.Request, db, hub, roomName, handleCommand)
	}
}

func WithWebSocketByDynamicRoomName(db *database.Manager, hub *ws.Hub, handleCommand func(c *ws.Client, cmd ws.Command)) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomName := c.Param("room")
		ws.HandleConnections(c.Writer, c.Request, db, hub, roomName, handleCommand)
	}
}
