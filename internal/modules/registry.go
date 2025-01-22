package module

import "go.uber.org/dig"

type Module interface {
	Name() string
	Provide(*dig.Container)
}

var registry = make(map[string]Module)

func Register(m Module) {
	registry[m.Name()] = m
}

func GetAll() []Module {
	modules := make([]Module, 0, len(registry))
	for _, m := range registry {
		modules = append(modules, m)
	}
	return modules
}
