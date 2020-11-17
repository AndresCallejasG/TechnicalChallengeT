package handlers

import (
	"TechnicalChallengeT/api/models"
	"encoding/csv"
	"net/http"
	"strconv"
)

//GetProductData extract data from an API and return a list of Products
func GetProductData(date string) ([]models.Product, error) {

	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"

	response, err := http.Get(url + "?date=" + date)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var prodList []models.Product

	csvProd := csv.NewReader(response.Body)
	csvProd.Comma = '\''

	csvProdData, err := csvProd.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, elem := range csvProdData {

		price, err := strconv.Atoi(elem[2])
		if err != nil {
			return nil, err
		}
		product, err := models.NewProduct(elem[0], elem[1], price)
		if err != nil {
			return nil, err
		}
		prodList = append(prodList, *product)
	}
	return prodList, nil
}
