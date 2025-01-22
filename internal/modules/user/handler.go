package user

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) (*Handler, error) {
	return &Handler{service: service}, nil
}

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
