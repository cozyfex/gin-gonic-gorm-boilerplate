package routing

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
	"gin-gonic-gorm-boilerplate/internal/ws"
)

func Route(r *gin.Engine, db *database.Manager, hub *ws.Hub) {
	RouteV1(r, db)
	RouteWS(r, db, hub)
}
