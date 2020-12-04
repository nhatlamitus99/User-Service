package main

import (
	"fmt"
	"net/http"

	"github.com/PhongVX/golang-rest-api/apis/userapi"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/{id}", userapi.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/users", userapi.ListUser).Methods("GET")
	router.HandleFunc("/api/v1/user", userapi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", userapi.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/v1/user/{id}", userapi.DeleteUser).Methods("DELETE")

	fmt.Println("Golang Rest API Is Running On Port: 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}

}
