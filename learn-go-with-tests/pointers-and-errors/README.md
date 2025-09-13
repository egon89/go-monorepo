# Pointers and Errors

- In Go, when you call a function or a method the arguments are copied.
- If you want to modify the value of an argument inside a function, you need to pass a pointer to the value instead of the value itself.

When we change the value of the balance inside the code, we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.
```go
func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	fmt.Printf("address of balance in test is %p and wallet %p\n", &wallet.balance, &wallet)

	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

# using non-pointer receiver
func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p and wallet %p\n", &w.balance, &w)
	w.balance += amount
}
```
Output:
```bash
=== RUN   TestWallet
address of balance in Deposit is 0xc0000121b0 and wallet 0xc0000121b0
address of balance in test is 0xc0000121a8 and wallet 0xc0000121a8
    wallet_test.go:20: got 0 want 10
--- FAIL: TestWallet (0.00s)
```

## nil
- `nil` is synonymous with `null` from other programming languages.
- Errors can be nil because the return type will be `error`, which is an *interface*.
  - a function that takes arguments or returns values that are *interfaces*, they can be nillable.
