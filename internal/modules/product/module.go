package product

import (
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	container.Provide(NewService)
	container.Provide(NewHandler)
	container.Provide(NewProductRoutes, dig.Group("routes"))
}
