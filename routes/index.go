package routes

import "github.com/gin-gonic/gin"

func AddRoutes(route *gin.RouterGroup) {
	eventRoutes(route)
}
