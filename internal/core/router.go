package core

import "github.com/gin-gonic/gin"

type Route interface {
	Register(router *gin.RouterGroup)
}

type RouteGroup struct {
	Prefix     string
	Middleware []gin.HandlerFunc
	Routes     []Route
}
