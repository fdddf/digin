//go:generate go run scripts/generate_modules.go
package main

import (
	"github.com/fdddf/digin/internal/core"
	_ "github.com/fdddf/digin/internal/gen"
)

func main() {
	container := core.BuildContainer()

	err := container.Invoke(func(server *core.Server) error {
		return server.Start()
	})

	if err != nil {
		panic(err)
	}
}
