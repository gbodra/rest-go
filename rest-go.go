package main

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFibonacci(w http.ResponseWriter, r *http.Request) {
	n := uint(10000)

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	json.NewEncoder(w).Encode(n2)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fibonacci", GetFibonacci).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
