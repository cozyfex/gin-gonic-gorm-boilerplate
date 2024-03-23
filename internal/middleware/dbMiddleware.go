package middleware

import (
	"gin-gonic-gorm-boilerplate/internal/database"
	"github.com/gin-gonic/gin"
)

func AddDbToContext(m *database.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbManger", m)
		c.Next()
	}
}
