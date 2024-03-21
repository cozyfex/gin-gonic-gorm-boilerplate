package routing

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	RouteV1(r)
}
