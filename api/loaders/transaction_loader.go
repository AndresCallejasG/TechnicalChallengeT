package loaders

import (
	"TechnicalChallengeT/api/handlers"
	"fmt"
)

//LoadTransactionData load a []Transaction to Dgraph
func LoadTransactionData(date string) error {

	transactions, err := handlers.GetTransactionData("")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(transactions)

	return nil
}
