//go:build ignore

// Solution for Exercise 06 — Mini Bank
// Run with: go run solution.go
package main

import (
	"errors"
	"fmt"
)

type InsufficientFundsError struct {
	Available float64
	Requested float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: have $%.2f, need $%.2f", e.Available, e.Requested)
}

type InvalidAmountError struct {
	Amount float64
}

func (e *InvalidAmountError) Error() string {
	return fmt.Sprintf("invalid amount: $%.2f (must be positive)", e.Amount)
}

type Transaction struct {
	Type    string
	Amount  float64
	Balance float64
}

type Account struct {
	Owner   string
	Balance float64
	History []Transaction
}

func NewAccount(owner string, initialBalance float64) (*Account, error) {
	if initialBalance < 0 {
		return nil, &InvalidAmountError{Amount: initialBalance}
	}
	a := &Account{Owner: owner, Balance: initialBalance}
	if initialBalance > 0 {
		a.History = append(a.History, Transaction{
			Type:    "initial deposit",
			Amount:  initialBalance,
			Balance: initialBalance,
		})
	}
	return a, nil
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return &InvalidAmountError{Amount: amount}
	}
	a.Balance += amount
	a.History = append(a.History, Transaction{
		Type:    "deposit",
		Amount:  amount,
		Balance: a.Balance,
	})
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return &InvalidAmountError{Amount: amount}
	}
	if amount > a.Balance {
		return &InsufficientFundsError{Available: a.Balance, Requested: amount}
	}
	a.Balance -= amount
	a.History = append(a.History, Transaction{
		Type:    "withdrawal",
		Amount:  amount,
		Balance: a.Balance,
	})
	return nil
}

func (a *Account) PrintStatement() {
	fmt.Printf("=== Statement: %s ===\n", a.Owner)
	for _, t := range a.History {
		fmt.Printf("  %-20s $%8.2f  → balance $%.2f\n", t.Type, t.Amount, t.Balance)
	}
	fmt.Printf("Current balance: $%.2f\n", a.Balance)
}

func main() {
	fmt.Println("=== Mini Bank ===\n")

	acc, err := NewAccount("Alice", 500.00)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	acc.Deposit(200.00)
	acc.Deposit(50.00)
	acc.Withdraw(100.00)

	fmt.Println("--- Attempting overdraft ---")
	if err := acc.Withdraw(10000.00); err != nil {
		fmt.Println("Error:", err)
		var insuf *InsufficientFundsError
		if errors.As(err, &insuf) {
			fmt.Printf("  → available: $%.2f, requested: $%.2f\n", insuf.Available, insuf.Requested)
		}
	}

	fmt.Println("\n--- Invalid deposit ---")
	if err := acc.Deposit(-50); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println()
	acc.PrintStatement()
}
