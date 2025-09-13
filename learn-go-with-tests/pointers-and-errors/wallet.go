package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin // private outside the package it's defined in
}

// so rather than taking a copy of the whole Wallet,
// we instead take a pointer to that wallet so that we can change the original values within it.
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p and wallet %p\n", &w.balance, &w)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// return (*w).balance  // without explicit dereference
	return w.balance // automatically dereferenced
}

// Stringer interface implementation
// Stringer is a built-in interface in the fmt package
// link: https://pkg.go.dev/fmt#Stringer
//
//	type Stringer interface {
//	    String() string
//	}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
