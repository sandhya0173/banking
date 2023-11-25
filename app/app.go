package app

import (
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", hello)
	http.HandleFunc("/customers", getAllCustomers)
	//starting server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
