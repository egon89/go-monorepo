package main

import "fmt"

type Wallet struct {
	balance int // private outside the package it's defined in
}

func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p and wallet %p\n", &w.balance, &w)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
