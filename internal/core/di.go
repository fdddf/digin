package core

import (
	"go.uber.org/dig"

	"github.com/fdddf/digin/internal/config"
	"github.com/fdddf/digin/internal/gen"
	"github.com/fdddf/digin/internal/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port   string
}

func BuildContainer() *dig.Container {
	container := dig.New()

	// Provide base config
	container.Provide(config.NewConfig)

	gen.AutoRegister(container)
	// Provide server
	container.Provide(NewServer)

	return container
}

type ServerParams struct {
	dig.In
	Config *config.Config
	Routes []routes.Route `group:"routes"`
}

func NewServer(params ServerParams) (*Server, error) {
	router := gin.Default()

	baseGroup := router.Group("/api/v1")
	for _, route := range params.Routes {
		route.Register(baseGroup)
	}

	return &Server{
		router: router,
		port:   params.Config.Port,
	}, nil
}

func (s *Server) Start() error {
	return s.router.Run(":" + s.port)
}
