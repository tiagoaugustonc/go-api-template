package di

import (
	"example.com/cadastro-power/pkg/adapter/out"
	"example.com/cadastro-power/pkg/application/person"
	"example.com/cadastro-power/pkg/infra/config"
	"example.com/cadastro-power/pkg/infra/orm"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {

	// config
	container.Provide(config.NewConfig)

	// Database
	container.Provide(orm.NewDB)

	// person
	container.Provide(out.NewPersonRepository)
	container.Provide(person.NewService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
