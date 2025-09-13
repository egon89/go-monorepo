package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	// the %p placeholder prints memory addresses in base 16 notation with leading 0xs
	fmt.Printf("address of balance in test is %p and wallet %p\n", &wallet.balance, &wallet)

	want := Bitcoin(10)

	message := fmt.Sprintf("got %s want %s", got, want)

	fmt.Println(message) // got 10 BTC want 10 BTC

	if got != want {
		t.Errorf(message)
	}
}
