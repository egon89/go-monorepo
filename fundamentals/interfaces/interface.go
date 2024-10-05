package interfaces

import "fmt"

type Person struct {
	Id   int
	Name string
}

type Dog struct {
	Id   int
	Name string
}

type Walker interface {
	Walk()
}

func (p Person) Walk() {
	fmt.Printf("person %v started walking\n", p.Name)
}

func (d Dog) Walk() {
	fmt.Printf("dog %v started walking\n", d.Name)
}

func InterfaceExample() {
	fmt.Println("> interfaceExample")
	person := Person{
		Id:   1,
		Name: "John Doe",
	}
	dog := Dog{
		Id:   90,
		Name: "Steve",
	}

	startWalking(person)
	startWalking(dog)
	/*
		ready to go!
		person John Doe started walking
		ready to go!
		dog Steve started walking
	*/
}

func startWalking(walker Walker) {
	fmt.Println("ready to go!")
	walker.Walk()
}
