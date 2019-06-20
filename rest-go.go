package main

import (
	// "fmt"
	// "time"
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"math/big"
)

func GetFibonacci(w http.ResponseWriter, r *http.Request) {
	// initTime := time.Now().Nanosecond()

	n:= uint(10000)
	
	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	// endTime := time.Now().Nanosecond()
	// fmt.Println("Time (nanosec): ", (endTime - initTime))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fibonacci", GetFibonacci).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
