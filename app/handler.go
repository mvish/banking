package app

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashi", "Delhi", "12345"},
		{"bob", "Boston", "747953"},
	}

	rw.Header().Add("Content-type", "application/json")
	json.NewEncoder(rw).Encode(customers)
}
