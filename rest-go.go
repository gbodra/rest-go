package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Price type holds the model for storing price values and
// exchanging them on the API
type Price struct {
	ID              string `json:"id,omitempty"`
	CustomerCluster string `json:"customercluster,omitempty"`
	Product         string `json:"product,omitempty"`
}

// CreateHash reads the json input and hashes it using HMAC SHA-256
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
	router.HandleFunc("/hash", CreateHash).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
