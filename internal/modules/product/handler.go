package product

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) (*Handler, error) {
	return &Handler{service: service}, nil
}

func (h *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.service.GetProduct(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
}
