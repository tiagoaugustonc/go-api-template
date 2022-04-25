package out

import (
	"fmt"

	"example.com/cadastro-power/pkg/application/person"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type memoryPersonRepository struct {
	gorm    *gorm.DB
	persons []person.Person
}

func NewPersonRepository(gorm *gorm.DB) person.Repository {
	return &memoryPersonRepository{gorm: gorm}
}

func (r *memoryPersonRepository) Create(person person.Person) (*person.Person, error) {
	fmt.Printf("Salvando pessoal %s\n", person)

	person.Id = uuid.NewString()

	r.persons = append(r.persons, person)

	return &person, nil
}

func (r *memoryPersonRepository) Retrieve(id string) (*person.Person, error) {

	for _, person := range r.persons {
		if person.Id == id {
			return &person, nil
		}
	}

	return nil, nil
}
