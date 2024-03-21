package routing

import (
	"gin-gonic-gorm-boilerplate/internal/api/v1"
	"github.com/gin-gonic/gin"
)

func RouteV1(r *gin.Engine) {
	v1Group := r.Group("v1")
	{
		v1Group.GET("/users", v1.GetUsers)
		v1Group.POST("/user", v1.PostUser)
	}
}
