package routing

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/ws"
	"gin-gonic-gorm-boilerplate/internal/ws/handle"
)

func RouteWS(r *gin.Engine, db *database.Manager, hub *ws.Hub) {
	wsGroup := r.Group("ws")
	{
		wsGroup.GET("/chat", WithWebSocket(db, hub, "chat", handle.ChatCommand))
		wsGroup.GET("/news", WithWebSocket(db, hub, "news", handle.NewsCommand))
		wsGroup.GET("/sports", WithWebSocket(db, hub, "sports", handle.SportsCommand))

		wsGroup.GET("/chat/:room", WithWebSocketByDynamicRoomName(db, hub, handle.ChatCommand))
	}
}
