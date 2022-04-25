package main

import (
	"fmt"

	"example.com/cadastro-power/pkg/adapter/out"
	"example.com/cadastro-power/pkg/application/person"
)

func main() {

	// Porta de saída da aplicação
	personRepository := out.NewPersonRepository(nil)

	// Service, porta de entrada para a aplicação
	personService := person.NewService(personRepository)

	// Usamos a porta de entrada para realizar um caso de uso. Criar uma pessoa
	personNew := person.Person{
		Fullname: "Fulano",
		Sex:      "M",
		Type:     "F",
		Document: "04481439688",
	}

	if personCreated, error := personService.Create(personNew); error != nil {
		fmt.Printf("Erro ao criar pessoa. %s\n", personNew)
	} else {
		fmt.Printf("Pessoa criada: %s\n", personCreated)
	}

}
