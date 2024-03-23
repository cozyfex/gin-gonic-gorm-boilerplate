package routing

import (
	"github.com/gin-gonic/gin"

	"gin-gonic-gorm-boilerplate/internal/api/v1"
	"gin-gonic-gorm-boilerplate/internal/database"
)

func RouteV1(r *gin.Engine, db *database.Manager) {
	v1Group := r.Group("v1")
	{
		v1Group.GET("/list-user", WithDB(db, v1.ListUser))
		v1Group.POST("/user", WithDB(db, v1.CreateUser))
	}
}
