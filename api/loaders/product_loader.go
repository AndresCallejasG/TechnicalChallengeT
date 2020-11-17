package loaders

import (
	"TechnicalChallengeT/api/handlers"
	"fmt"
)

//LoadProductData load a []Product to Dgraph
func LoadProductData(date string) error {

	products, err := handlers.GetProductData("")
	if err != nil {
		return err
	}

	fmt.Println(products)

	return nil
}
