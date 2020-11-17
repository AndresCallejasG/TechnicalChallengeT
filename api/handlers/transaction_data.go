package handlers

import (
	"TechnicalChallengeT/api/models"
	"bytes"
	"net/http"
	"strings"
)

//GetTransactionData extract data from an API and return a list of transactions
func GetTransactionData(date string) ([]models.Transaction, error) {

	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"

	response, err := http.Get(url + "?date=" + date)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var tranList []models.Transaction

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	newStr := buf.String()

	splitData := strings.Split(newStr, "#")

	//This function could be improve
	for _, elem := range splitData {
		splitTran := strings.Split(elem, "\x00")
		if len(splitTran) > 4 {
			prodStr := splitTran[4][1 : len(splitTran[4])-1]
			splitProd := strings.Split(prodStr, ",")
			tran, err := models.NewTransaction(splitTran[0], splitTran[1], splitTran[2], splitTran[3], splitProd)
			if err != nil {
				return nil, err
			}
			tranList = append(tranList, *tran)
		}
	}

	return tranList, nil
}
