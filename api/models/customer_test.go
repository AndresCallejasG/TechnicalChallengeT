package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCustomer(t *testing.T) {
	c := require.New(t)

	cust, err := NewCustomer("231d23b6", "Estell", 32)

	c.NotEmpty(cust.CustomerID)
	c.NotEmpty(cust.Name)

	c.Equal("231d23b6", cust.CustomerID)
	c.Equal("Estell", cust.Name)
	c.Equal(32, cust.Age)

	c.NoError(err)
}

func TestCustomerErr(t *testing.T) {
	c := require.New(t)

	_, err := NewCustomer("", "Estell", 32)
	c.Error(ErrEmptyCustID, err)

	_, err = NewCustomer("231d23b6", "", 32)
	c.Error(ErrEmptyCustName, err)

	_, err = NewCustomer("231d23b6", "Estell", -1)
	c.Error(ErrInvalidCustAge, err)

}

func BenchmarkNewCustomer(b *testing.B) {
	c := require.New(b)

	for n := 0; n < b.N; n++ {
		_, err := NewCustomer("231d23b6", "Estell", 32)
		c.NoError(err)
	}
}
