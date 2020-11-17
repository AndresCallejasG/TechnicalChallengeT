package main

import (
	"fmt"
	"net/http"

	"TechnicalChallengeT/api/loaders"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	port := ":3000"

	r.Use(middleware.Logger)
	r.Get("/v1/load", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, make a POST and send date (unix timestamp) in the body to load data to Dgraph"))
	})
	r.Post("/v1/load", loadData)

	fmt.Println("ETL Server running in localhost" + port)
	fmt.Println("load data using the endpoint localhost" + port + "/v1/load")
	fmt.Println("...")
	http.ListenAndServe(port, r)
}

func loadData(w http.ResponseWriter, r *http.Request) {

	err := loaders.LoadCustomerData("")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("-----------------------------------")

	err = loaders.LoadProductData("")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("-----------------------------------")

	err = loaders.LoadTransactionData("")
	if err != nil {
		fmt.Print(err)
	}

	w.Write([]byte("Data fetched successfully"))
}
