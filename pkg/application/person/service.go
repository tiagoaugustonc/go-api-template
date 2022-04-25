package person

import "fmt"

type Service interface {
	Create(person Person) (*Person, error)
	Retrieve(id string) (*Person, error)
	Update(person Person) (*Person, error)
	Delete(id string) (bool, error)
}

type personService struct {
	repo Repository
}

func NewService(personRepository Repository) Service {

	return &personService{
		repo: personRepository,
	}
}

func (p *personService) Create(person Person) (*Person, error) {
	fmt.Printf("Criando pessoa %s\n", person)

	createdPerson, _ := p.repo.Create(person)

	return createdPerson, nil
}

func (p *personService) Retrieve(id string) (*Person, error) {
	fmt.Printf("Retriving person %s\n", p)

	return p.repo.Retrieve(id)
}

func (p *personService) Update(person Person) (*Person, error) {
	fmt.Printf("Updating person %s\n", p)
	return &person, nil
}

func (p *personService) Delete(id string) (bool, error) {
	fmt.Printf("Deleting person %s", p)
	return true, nil
}
