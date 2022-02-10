package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:4000", mux))
}
