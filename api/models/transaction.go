package models

import (
	"errors"
)

var (
	// ErrEmptyTranID if the TransactionID is empty.
	ErrEmptyTranID = errors.New("'transactionID' cannot be empty")

	// ErrEmptyTranBuyer if the Buyer is empty.
	ErrEmptyTranBuyer = errors.New("'buyer' cannot be empty")

	// ErrEmptyTranDevice if the Device is a negative int
	ErrEmptyTranDevice = errors.New("'device' cannot be empty")

	// ErrEmptyTranIP if the ip is a negative int
	ErrEmptyTranIP = errors.New("'ip' cannot be empty")

	// ErrInvalidTranProducts if products is null or empty
	ErrInvalidTranProducts = errors.New("'products' must have at list 1 product")
)

// Transaction struct to represent a Transaction
type Transaction struct {
	TransactionID string   `json:"id"`
	Buyer         string   `json:"buyer"`
	Device        string   `json:"device"`
	IP            string   `json:"ip"`
	Products      []string `json:"products"`
}

// NewTransaction constructor
func NewTransaction(transactionID, buyer, device, ip string, products []string) (*Transaction, error) {
	if transactionID == "" {
		return nil, ErrEmptyTranID
	}

	if buyer == "" {
		return nil, ErrEmptyTranBuyer
	}

	if device == "" {
		return nil, ErrEmptyTranDevice
	}

	if ip == "" {
		return nil, ErrEmptyTranIP
	}

	if products == nil || len(products) == 0 {
		return nil, ErrInvalidTranProducts
	}

	newTransaction := &Transaction{
		TransactionID: transactionID,
		Buyer:         buyer,
		Device:        device,
		IP:            ip,
		Products:      products,
	}

	return newTransaction, nil
}
