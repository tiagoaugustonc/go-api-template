package person

type Repository interface {
	Create(person Person) (*Person, error)
	Retrieve(id string) (*Person, error)
}
