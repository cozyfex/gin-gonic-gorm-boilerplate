package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 출처 검사 비활성화 (보안 상 생산 환경에서는 수정 필요)
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request, db *database.Manager, hub *Hub, roomName string, handleCommand func(*Client, Command)) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error(fmt.Sprintln("Failed to upgrade to websocket:", err))
		return
	}
	if _, exists := hub.Rooms[roomName]; !exists {
		hub.Rooms[roomName] = NewRoom(roomName)
	}
	client := NewClient(ws, db, hub.Rooms[roomName], handleCommand)
	client.Room.Register <- client

	go client.readPump()
	go client.writePump()
}
