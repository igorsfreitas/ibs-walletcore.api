package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAccount creates a new account
func NewAccount(client *Client) *Account {
	if client == nil {
		return nil
	}
	return &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Deposit deposits the amount into the account
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	a.UpdatedAt = time.Now()
}

// Withdraw withdraws the amount from the account
func (a *Account) Withdraw(amount float64) {
	a.Balance -= amount
	a.UpdatedAt = time.Now()
}
