package loaders

import (
	"TechnicalChallengeT/api/handlers"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//LoadTransactionData load a []Transaction to Dgraph
func LoadTransactionData(date string) error {

	transactions, err := handlers.GetTransactionData(date)
	if err != nil {
		fmt.Print(err)
	}

	url := "http://localhost:8080/graphql"
	method := "POST"

	var tranStr string
	tranStr = ""

	for i, elem := range transactions {
		if i != 0 {
			tranStr += ","
		}
		tranStr += elem.TransactionInputStr()

		// Limit that should be controled better -> problems with my pc performance
		if i == 200 {
			break
		}
	}

	mutString := "{\"query\":\"mutation addTransaction($transactions: [AddTransactionInput!]!){\\n  addTransaction(input: $transactions){\\n    transaction {\\n      transactionID\\n      ip\\n      device\\n    }\\n  }\\n}\",\"variables\":{\"transactions\":["
	mutString += tranStr + "]}}"

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

	return nil
}
