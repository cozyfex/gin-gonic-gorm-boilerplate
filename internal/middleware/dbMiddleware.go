package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/database"
)

func AddDbToContext(m *database.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbManger", m)
		c.Next()
	}
}
