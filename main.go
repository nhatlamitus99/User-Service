package main

import (
	"fmt"
	"net/http"

	"github.com/PhongVX/golang-rest-api/apis/userapi"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/authorize", userapi.Authorize).Methods("POST")
	router.HandleFunc("/api/resource", userapi.GetResource).Methods("GET")

	fmt.Println("Golang Rest API Is Running On Port: 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}

}
