package person

import "fmt"

type Talker interface {
	talking()
}

type Person struct {
	name string
}

func (p *Person) talking() {
	fmt.Println("hai")
}

func NewPerson () *Person{
	return &Person{
		name: "cris",
	}
}

func NewPersonTalker() Talker {
	return &Person{
		name: "desman",
	}
}

func (person *Person) SetPersonName(name string){
	person.name = name
}

func (person *Person) GetPersonName() string {
	return person.name
}