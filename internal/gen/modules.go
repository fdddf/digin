// Code generated by go generate; DO NOT EDIT.
package gen

import (
	"go.uber.org/dig"

	"github.com/fdddf/digin/internal/modules/product"
	"github.com/fdddf/digin/internal/modules/user"
)

func AutoRegister(container *dig.Container) {
	product.Register(container)
	user.Register(container)
}
