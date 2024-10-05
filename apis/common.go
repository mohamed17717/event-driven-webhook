package apis

import (
	"event-driven-webhook/config"
	"event-driven-webhook/middlewares"
	"github.com/gin-gonic/gin"
)

func ProtectedRoute() gin.RouterGroup {
	var route = config.Server.Group("/")
	route.Use(middlewares.IsAuthenticated())

	return *route
}

type Empty struct {
}
