package routing

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
)

func WithDB(db *database.Manager, h func(c *gin.Context, db *database.Manager)) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c, db)
	}
}
