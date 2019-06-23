package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Price struct {
	ID              string `json:"id,omitempty"`
	CustomerCluster string `json:customercluster,omitempty`
	Product         string `json:product,omitempty`
}

func GetFibonacci(w http.ResponseWriter, r *http.Request) {
	n := uint(10000)

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}

	json.NewEncoder(w).Encode(n2)
}

func CreateHash(w http.ResponseWriter, r *http.Request) {
	bodyData, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	price := string(bodyData)

	secret := "mysecret"
	fmt.Printf("Secret: %s Data: %s\n", secret, price)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(price))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	json.NewEncoder(w).Encode(sha)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fibonacci", GetFibonacci).Methods("GET")
	router.HandleFunc("/hash", CreateHash).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
