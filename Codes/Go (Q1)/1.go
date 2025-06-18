package main

import (
	"fmt"
)

// BankAccount struct
type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit method: adds amount to balance
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("%s deposited: $%.2f\n", b.Owner, amount)
}

// Withdraw method: subtracts if enough balance
func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.Balance {
		fmt.Printf(" Withdrawal of $%.2f failed: Insufficient funds!\n", amount)
	} else {
		b.Balance -= amount
		fmt.Printf(" %s withdrew: $%.2f\n", b.Owner, amount)
	}
}

// BalanceInquiry method: shows current balance
func (b BankAccount) BalanceInquiry() {
	fmt.Printf(" Current balance for %s: $%.2f\n", b.Owner, b.Balance)
}

func main() {
	// Create a bank account
	account := BankAccount{Owner: "Ahad"}

	// 3 demo transactions
	account.Deposit(1000)
	account.Withdraw(250)
	account.Withdraw(1000) // This should fail

	// Final balance check
	account.BalanceInquiry()
}
