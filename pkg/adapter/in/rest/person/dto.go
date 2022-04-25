package person

import "example.com/cadastro-power/pkg/application/person"

type PersonDTO struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname" binding:"required"`
	Sex      string `json:"sex" binding:"required,oneof=M F"`
	Type     string `json:"type" binding:"required,oneof=F J"`
	Document string `json:"document" binding:"required"`
}

func (d *PersonDTO) toModel() *person.Person {

	model := person.Person{
		Id:       d.Id,
		Fullname: d.Fullname,
		Sex:      d.Sex,
		Document: d.Document,
	}

	return &model
}
