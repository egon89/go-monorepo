package main

import "fmt"

type Wallet struct {
	balance int // private outside the package it's defined in
}

// so rather than taking a copy of the whole Wallet,
// we instead take a pointer to that wallet so that we can change the original values within it.
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p and wallet %p\n", &w.balance, &w)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// return (*w).balance  // without explicit dereference
	return w.balance // automatically dereferenced
}
