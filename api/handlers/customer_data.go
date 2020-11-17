package handlers

import (
	"TechnicalChallengeT/api/models"
	"encoding/json"
	"net/http"
)

//GetCustomerData extract data from an API and return a list of Buyers
func GetCustomerData(date string) ([]models.Customer, error) {

	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"

	response, err := http.Get(url + "?date=" + date)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var custList []models.Customer

	json.NewDecoder(response.Body).Decode(&custList)

	return custList, nil
}
