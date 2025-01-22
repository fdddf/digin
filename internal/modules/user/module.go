package user

import (
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	container.Provide(NewService)
	container.Provide(NewHandler)
	container.Provide(NewUserRoutes, dig.Group("routes"))
}
