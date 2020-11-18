package loaders

import (
	"TechnicalChallengeT/api/handlers"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//LoadCustomerData load []Customer to Dgraph
func LoadCustomerData(date string) error {

	//var customers []models.Customer

	customers, err := handlers.GetCustomerData(date)
	if err != nil {
		return err
	}

	url := "http://localhost:8080/graphql"
	method := "POST"

	var custStr string
	custStr = ""

	for i, elem := range customers {
		if i != 0 {
			custStr += ","
		}
		custStr += elem.CustomerInputStr()
	}

	mutString := "{\"query\":\"mutation addCustomer($customers: [AddCustomerInput!]!){\\n  addCustomer(input: $customers){\\n    customer {\\n      customerID\\n      name\\n      age\\n    }\\n  }\\n}\",\"variables\":{\"customers\":["
	mutString += custStr + "]}}"

	//fmt.Println(mutString)

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

	//fmt.Println(customers)

	return nil
}
