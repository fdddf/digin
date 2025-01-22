package product

import (
	"github.com/fdddf/digin/internal/routes"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	handler *Handler
}

func NewProductRoutes(handler *Handler) routes.Route {
	return &ProductRoutes{handler: handler}
}

func (r *ProductRoutes) Register(group *gin.RouterGroup) {
	productGroup := group.Group("/products")
	{
		productGroup.GET("/:id", r.handler.GetProduct)
		// productGroup.POST("", r.handler.CreateProduct)
		// Add more product-related routes
	}
}
