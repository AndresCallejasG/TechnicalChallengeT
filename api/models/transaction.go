package models

import (
	"errors"
	"fmt"
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

	// ErrInvalidTranProducts if Products is null or empty
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
func NewTransaction(transactionID, buyer, ip, device string, products []string) (*Transaction, error) {
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

//TransactionInputStr return the str representation to input in GraphQL
func (tran *Transaction) TransactionInputStr() string {

	var str string

	str = "{\"transactionID\":\"" + tran.TransactionID +
		"\",\"buyer\":{\"customerID\":\"" + tran.Buyer +
		"\"},\"ip\":\"" + tran.IP + "\",\"device\":\"" +
		tran.Device + "\",\"products\":["

	for i, elem := range tran.Products {
		if i != 0 {
			str += ","
		}
		str = str + "{\"productID\":\"" + elem + "\"}"
	}

	str = str + "]}"

	fmt.Println(str)
	return str
}
