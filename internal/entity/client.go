package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewClient creates a new client
func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := client.Validate(); err != nil {
		return nil, err
	}

	return client, nil
}

// Validate validates the client
func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

// Update updates the client
func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}

// AddAccount adds an account to the client
func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("account does not belong to this client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
