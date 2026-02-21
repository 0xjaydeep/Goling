// Exercise 06 — Mini Bank
//
// Problem:
//   Build a banking system with proper error handling using custom error types.
//
// Types to implement:
//   type InsufficientFundsError struct { Available, Requested float64 }
//   type InvalidAmountError struct { Amount float64 }
//   both implement the error interface
//
//   type Transaction struct { Type string; Amount, Balance float64 }
//   type Account struct { Owner string; Balance float64; History []Transaction }
//
// Functions to implement:
//   NewAccount(owner string, initialBalance float64) (*Account, error)
//     — error (InvalidAmountError) if initialBalance < 0
//   (a *Account) Deposit(amount float64) error
//     — error (InvalidAmountError) if amount <= 0
//   (a *Account) Withdraw(amount float64) error
//     — error (InvalidAmountError) if amount <= 0
//     — error (InsufficientFundsError) if amount > a.Balance
//   (a *Account) PrintStatement()
//     — prints all transactions with running balance
//
// Run your solution:  go run main.go
// See the answer:     go run solution.go
package main

import (
	"errors"
	"fmt"
)

// InsufficientFundsError is returned when a withdrawal exceeds the current balance.
type InsufficientFundsError struct {
	Available float64
	Requested float64
}

func (e *InsufficientFundsError) Error() string {
	// TODO: return fmt.Sprintf("insufficient funds: have $%.2f, need $%.2f", e.Available, e.Requested)
	return "not implemented"
}

// InvalidAmountError is returned when a deposit or withdrawal amount is <= 0.
type InvalidAmountError struct {
	Amount float64
}

func (e *InvalidAmountError) Error() string {
	// TODO: return fmt.Sprintf("invalid amount: $%.2f (must be positive)", e.Amount)
	return "not implemented"
}

// Transaction records a single account event.
type Transaction struct {
	Type    string  // "initial deposit", "deposit", or "withdrawal"
	Amount  float64
	Balance float64 // balance AFTER this transaction
}

// Account represents a bank account.
type Account struct {
	Owner   string
	Balance float64
	History []Transaction
}

// NewAccount creates a new Account.
// Returns an error if initialBalance is negative.
func NewAccount(owner string, initialBalance float64) (*Account, error) {
	// TODO: if initialBalance < 0, return nil and &InvalidAmountError{Amount: initialBalance}
	// TODO: create the Account
	// TODO: if initialBalance > 0, append an "initial deposit" Transaction to History
	// TODO: return the account and nil
	return nil, fmt.Errorf("not implemented")
}

// Deposit adds amount to the account.
// Returns InvalidAmountError if amount <= 0.
func (a *Account) Deposit(amount float64) error {
	// TODO: validate amount > 0
	// TODO: a.Balance += amount
	// TODO: append Transaction{Type: "deposit", Amount: amount, Balance: a.Balance}
	// TODO: return nil
	return fmt.Errorf("not implemented")
}

// Withdraw subtracts amount from the account.
// Returns InvalidAmountError if amount <= 0.
// Returns InsufficientFundsError if amount > a.Balance.
func (a *Account) Withdraw(amount float64) error {
	// TODO: validate amount > 0
	// TODO: check a.Balance >= amount
	// TODO: a.Balance -= amount
	// TODO: append Transaction{Type: "withdrawal", Amount: amount, Balance: a.Balance}
	// TODO: return nil
	return fmt.Errorf("not implemented")
}

// PrintStatement prints all transactions and the current balance.
func (a *Account) PrintStatement() {
	// TODO: fmt.Printf("=== Statement: %s ===\n", a.Owner)
	// TODO: for each transaction:
	//         fmt.Printf("  %-20s $%8.2f  → balance $%.2f\n", t.Type, t.Amount, t.Balance)
	// TODO: fmt.Printf("Current balance: $%.2f\n", a.Balance)
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

	// Overdraft attempt
	fmt.Println("--- Attempting overdraft ---")
	if err := acc.Withdraw(10000.00); err != nil {
		fmt.Println("Error:", err)
		var insuf *InsufficientFundsError
		if errors.As(err, &insuf) {
			fmt.Printf("  → available: $%.2f, requested: $%.2f\n", insuf.Available, insuf.Requested)
		}
	}

	// Invalid amount
	fmt.Println("\n--- Invalid deposit ---")
	if err := acc.Deposit(-50); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println()
	acc.PrintStatement()
}
