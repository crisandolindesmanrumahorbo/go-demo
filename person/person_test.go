package person

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	person := Person{name: "cris"}
	modify(&person)
	modifyy(person)
	fmt.Println("person", person.name)
}

func modify(person *Person) {
	person.name = "change"
}

func modifyy(person Person) {
	person.name = "it"
	fmt.Println("person moduf", person.name)
}

func TestNewPerson(t *testing.T) {
	var crisPerson *Person
	crisPerson = NewPerson()
	crisPerson.SetPersonName("crisan")
	fmt.Println(crisPerson.GetPersonName())
}

func TestInteface(t *testing.T) {
	// personTalker := NewPersonTalker()
	// personTalker.talking()

	var personTalker Talker
	ref := &Person{name: "cris"}
	personTalker = ref
	personTalker.talking()

	var personTalker2 Talker
	val := Person{name: "des"}
	personTalker2 = &val
	personTalker2.talking()
}