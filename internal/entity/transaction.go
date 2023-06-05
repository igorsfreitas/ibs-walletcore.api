package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

// NewTransaction creates a new transaction
func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	if accountFrom == nil || accountTo == nil {
		return nil, errors.New("account is required")
	}

	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	transaction.Commit()

	return transaction, nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Withdraw(t.Amount)
	t.AccountTo.Deposit(t.Amount)
}

// Validate validates the transaction
func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}

	return nil
}
