package middleware

import (
	"gin-gonic-gorm-boilerplate/internal/db"
	"github.com/gin-gonic/gin"
)

func AddDbToContext(m *db.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbManger", m)
		c.Next()
	}
}
