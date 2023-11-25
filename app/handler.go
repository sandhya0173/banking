package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type customers struct {
	Name    string `json:"full_name" xml:"Fname"`
	City    string `json:"city" xml:"Lname"`
	Zipcode int    `json:"zipcode_" xml:"zcode"`
}

// define routea
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	Customers := []customers{
		{"ashish", "Delhi", 1234},
		{"john", "minosatta", 6589},
	}
	if r.Header.Get("content-Type") == "application/xml" {
		w.Header().Add("content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(Customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		w.Header().Add("content-Type", "application/json")
		err := json.NewEncoder(w).Encode(Customers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
