package loaders

import (
	"TechnicalChallengeT/api/handlers"
	"fmt"
)

//LoadCustomerData load []Customer to Dgraph
func LoadCustomerData(date string) error {

	customers, err := handlers.GetCustomerData("")
	if err != nil {
		return err
	}

	fmt.Println(customers)

	return nil
}
