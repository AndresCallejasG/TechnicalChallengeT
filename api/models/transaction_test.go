package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	c := require.New(t)

	tran, err := NewTransaction("00005fb1c100", "76189a41", "92.69.220.153", "mac", []string{"12e568e0", "9a2841a3"})

	c.NotEmpty(tran.TransactionID)
	c.NotEmpty(tran.Buyer)
	c.NotEmpty(tran.Device)
	c.NotEmpty(tran.IP)

	c.Equal("00005fb1c100", tran.TransactionID)
	c.Equal("76189a41", tran.Buyer)
	c.Equal("92.69.220.153", tran.Device)
	c.Equal("mac", tran.IP)
	c.Equal([]string{"12e568e0", "9a2841a3"}, tran.Products)

	c.NoError(err)
}

func TestTransactionErr(t *testing.T) {
	c := require.New(t)

	_, err := NewTransaction("", "76189a41", "92.69.220.153", "mac", []string{"12e568e0", "9a2841a3"})
	c.Error(ErrEmptyTranID, err)

	_, err = NewTransaction("00005fb1c100", "", "92.69.220.153", "mac", []string{"12e568e0", "9a2841a3"})
	c.Error(ErrEmptyTranBuyer, err)

	_, err = NewTransaction("00005fb1c100", "76189a41", "", "mac", []string{"12e568e0", "9a2841a3"})
	c.Error(ErrEmptyTranIP, err)

	_, err = NewTransaction("00005fb1c100", "76189a41", "92.69.220.153", "", []string{"12e568e0", "9a2841a3"})
	c.Error(ErrEmptyTranDevice, err)

	_, err = NewTransaction("00005fb1c100", "76189a41", "92.69.220.153", "mac", []string{})
	c.Error(ErrInvalidTranProducts, err)
}

func BenchmarkNewTransaction(b *testing.B) {
	c := require.New(b)

	for n := 0; n < b.N; n++ {
		_, err := NewTransaction("00005fb1c100", "76189a41", "92.69.220.153", "mac", []string{"12e568e0", "9a2841a3"})
		c.NoError(err)
	}
}
