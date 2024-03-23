package routing

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
)

func Route(r *gin.Engine, db *database.Manager) {
	RouteV1(r, db)
}
