package pointers

import "fmt"

type Account struct {
	balance float64
}

func PointerExample() {
	fmt.Println("> pointerExample")
	a := 10
	fmt.Printf("a: %v\n", a)
	// * (asterisk) -> dereference
	var aPointer *int = &a
	fmt.Printf("aPointer: %v\n", aPointer)

	*aPointer = 20
	fmt.Printf("a: %v\n", a)
	b := &a
	fmt.Printf("b as pointer: %v, b as pointer value using *: %v\n", b, *b)

	*b = 90
	fmt.Printf("a: %v, *aPointer: %v, *b: %v\n", a, *aPointer, *b)
	fmt.Printf("&a: %v, aPointer: %v, b: %v\n", &a, aPointer, b)

	/*
		a: 10
		aPointer: 0xc00009a1d0
		a: 20
		b as pointer: 0xc00009a1d0, b as pointer value using *: 20
		a: 90, *aPointer: 90, *b: 90
		&a: 0xc00009a1d0, aPointer: 0xc00009a1d0, b: 0xc00009a1d0
	*/

	account := Account{
		balance: 100.0,
	}
	fmt.Printf("addWithoutChange: %v\n", account.addWithoutChange(50))
	fmt.Printf("account.total: %v\n", account.balance)

	fmt.Printf("add: %v\n", account.add(50))
	fmt.Printf("account.total: %v\n", account.balance)

	newAccount := createAccount(20)
	fmt.Printf("create account: %v\n", newAccount.balance)

	/*
		addWithoutChange: 150
		account.total: 100
		add: 150
		account.total: 150
		create account: 20
	*/
}

func (a Account) addWithoutChange(value float64) float64 {
	// we are working with a copy of the account
	a.balance += value

	return a.balance
}

func (a *Account) add(value float64) float64 {
	// we are working with the account reference
	a.balance += value

	return a.balance
}

func createAccount(balance float64) *Account {
	return &Account{balance}
}
