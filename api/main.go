package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	dateUnix := r.URL.Query().Get("date")

	if dateUnix != "" {
		const layout = "2006-01-02"

		t, err := time.Parse(layout, dateUnix)
		if err != nil {
			log.Fatal(err)
		}

		dateUnix = strconv.Itoa(int(t.Unix()))
	}

	// fmt.Println(dateUnix)

	err := loaders.LoadCustomerData(dateUnix)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("-------------Customers Loaded----------------------")

	err = loaders.LoadProductData(dateUnix)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("----------------Products Loaded-------------------")

	err = loaders.LoadTransactionData(dateUnix)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("----------------Transaction Loaded-------------------")

	w.Write([]byte("Data fetched successfully"))
}
