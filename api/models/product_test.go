package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct(t *testing.T) {
	c := require.New(t)

	prod, err := NewProduct("32af6187", "White rice", 4788)

	c.NotEmpty(prod.ProductID)
	c.NotEmpty(prod.Name)

	c.Equal("32af6187", prod.ProductID)
	c.Equal("White rice", prod.Name)
	c.Equal(4788, prod.Price)

	c.NoError(err)
}

func TestProductErr(t *testing.T) {
	c := require.New(t)

	_, err := NewProduct("", "White rice", 4788)
	c.Error(ErrEmptyProdID, err)

	_, err = NewProduct("32af6187", "", 4788)
	c.Error(ErrEmptyProdName, err)

	_, err = NewProduct("32af6187", "White rice", -1)
	c.Error(ErrInvalidProdPrice, err)

}

func BenchmarkNewProduct(b *testing.B) {
	c := require.New(b)
	for n := 0; n < b.N; n++ {
		_, err := NewProduct("32af6187", "White rice", 4788)
		c.NoError(err)
	}
}
