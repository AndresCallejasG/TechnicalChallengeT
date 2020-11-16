package models

import (
	"errors"
)

var (
	// ErrEmptyCustID if the customerID is empty.
	ErrEmptyCustID = errors.New("'customerID' cannot be empty")

	// ErrEmptyCustName if the Name is empty.
	ErrEmptyCustName = errors.New("'name' cannot be empty")

	// ErrInvalidCustAge if the age is a negative int
	ErrInvalidCustAge = errors.New("'age' cannot be negative")
)

// Customer struct to represent a customer
type Customer struct {
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
}

// NewCustomer constructor
func NewCustomer(customerID, name string, age int) (*Customer, error) {
	if customerID == "" {
		return nil, ErrEmptyCustID
	}

	if name == "" {
		return nil, ErrEmptyCustName
	}

	if age < 0 {
		return nil, ErrInvalidCustAge
	}

	newCustomer := &Customer{
		CustomerID: customerID,
		Name:       name,
		Age:        age,
	}

	return newCustomer, nil
}
