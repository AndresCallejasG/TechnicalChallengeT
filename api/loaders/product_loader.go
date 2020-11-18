package loaders

import (
	"TechnicalChallengeT/api/handlers"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//LoadProductData load a []Product to Dgraph
func LoadProductData(date string) error {

	products, err := handlers.GetProductData(date)
	if err != nil {
		return err
	}

	url := "http://localhost:8080/graphql"
	method := "POST"

	var prodStr string
	prodStr = ""

	for i, elem := range products {
		if i != 0 {
			prodStr += ","
		}
		prodStr += elem.ProductInputStr()
	}

	mutString := "{\"query\":\"mutation addProduct($products: [AddProductInput!]!){\\n  addProduct(input: $products){\\n    product {\\n      productID\\n      name\\n      price\\n    }\\n  }\\n}\",\"variables\":{\"products\":["
	mutString += prodStr + "]}}"

	payload := strings.NewReader(mutString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println(string(body))

	// mutatation to save to Dgraph

	//fmt.Println(products)

	return nil
}
