package structs

import (
	"fmt"
	"time"
)

type Contact struct {
	Phone string
	Email string
}

type Address struct {
	Street string
	Number int
}

type Client struct {
	Document  string
	Name      string
	CreatedAt time.Time
	Contact
	Address Address
	Active  bool
}

func StructsExample() {
	fmt.Println("> structsExample")
	user := Client{
		Document:  "00011122233",
		Name:      "John Doe",
		CreatedAt: time.Now(),
		Address: Address{
			Street: "Street A",
		},
		Active: true,
	}
	user.Phone = "000000000"
	user.Email = "johndoe@domain.com"
	user.Address.Number = 10

	user.print()
	// document: 00011122233, name: John Doe, email: johndoe@domain.com, street: Street A
}

func (c Client) print() {
	fmt.Printf("document: %v, name: %v, email: %v, street: %v\n", c.Document, c.Name, c.Email, c.Address.Street)
}
