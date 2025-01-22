package routes

import "github.com/gin-gonic/gin"

// Route defines the contract for route registration
type Route interface {
	Register(router *gin.RouterGroup)
}

// RouteGroup configuration
type RouteGroup struct {
	Prefix     string
	Middleware []gin.HandlerFunc
	Routes     []Route
}
