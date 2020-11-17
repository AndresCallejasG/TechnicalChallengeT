package models

import (
	"errors"
)

var (
	// ErrEmptyProdID if the ProductID is empty.
	ErrEmptyProdID = errors.New("'ProductID' cannot be empty")

	// ErrEmptyProdName if the Name is empty.
	ErrEmptyProdName = errors.New("'name' cannot be empty")

	// ErrInvalidProdPrice if the Price is a negative int
	ErrInvalidProdPrice = errors.New("'price' cannot be negative")
)

// Product struct to represent a Product
type Product struct {
	ProductID string `json:"id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
}

// NewProduct constructor
func NewProduct(productID, name string, price int) (*Product, error) {
	if productID == "" {
		return nil, ErrEmptyProdID
	}

	if name == "" {
		return nil, ErrEmptyProdName
	}

	if price < 0 {
		return nil, ErrInvalidProdPrice
	}

	newProduct := &Product{
		ProductID: productID,
		Name:      name,
		Price:     price,
	}

	return newProduct, nil
}
