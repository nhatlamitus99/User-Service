package main

import (
	"GolangRestApi/apis/userapi"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", userapi.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", userapi.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", userapi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", userapi.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", userapi.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		panic(err)
	}
}
