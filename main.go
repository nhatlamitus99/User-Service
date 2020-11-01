package main

import (
	"fmt"
	"net/http"

	"github.com/PhongVX/golang-rest-api/apis/userapi"
	"github.com/PhongVX/golang-rest-api/db"
	"github.com/gorilla/mux"
)

func main() {
	db.GetDB()
	router := mux.NewRouter()

	router.HandleFunc("/authorize", userapi.Authorize).Methods("POST")
	router.HandleFunc("/api/userInfo", userapi.GetInfo).Methods("GET")

	fmt.Print("Golang Rest API Is Running On Port: 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}

}
