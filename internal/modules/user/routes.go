package user

import (
	"github.com/fdddf/digin/internal/routes" // Use routes package instead of core
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *Handler
}

func NewUserRoutes(handler *Handler) routes.Route { // Return routes.Route interface
	return &UserRoutes{handler: handler}
}

func (r *UserRoutes) Register(group *gin.RouterGroup) {
	userGroup := group.Group("/user")
	{
		userGroup.GET("/:id", r.handler.GetUser)
		// userGroup.POST("", r.handler.CreateUser)
	}
}
