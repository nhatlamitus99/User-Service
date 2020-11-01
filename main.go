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
	router.HandleFunc("/api/resource", userapi.GetResource).Methods("GET")
	router.HandleFunc("/api/v1/user/find", userapi.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/create", userapi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", userapi.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", userapi.Delete).Methods("DELETE")

	fmt.Println("Golang Rest API Is Running On Port: 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}

}
