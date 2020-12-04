package userapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/PhongVX/golang-rest-api/db"
	"github.com/PhongVX/golang-rest-api/entities"
)

func GetUser(response http.ResponseWriter, request *http.Request) {

	id := request.Header.Get("id")
	idUser, _ := strconv.Atoi(id)
	user := db.GetUser(idUser)

	responseWithJSON(response, http.StatusOK, user)

}

func ListUser(response http.ResponseWriter, request *http.Request) {

	id := request.Header.Get("id")
	idUser, _ := strconv.Atoi(id)
	user := db.GetUser(idUser)

	responseWithJSON(response, http.StatusOK, user)

}

func CreateUser(response http.ResponseWriter, request *http.Request) {

	user := entities.User{}
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	}

	err = db.CreateUser(user)
	if err != nil {
		responseWithError(response, http.StatusForbidden, "forbidden")
	}
	responseWithJSON(response, http.StatusOK, user)

}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	user := entities.User{}
	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	}

	err = db.UpdateUser(user)
	if err != nil {
		responseWithError(response, http.StatusForbidden, "forbidden")
	}
	responseWithJSON(response, http.StatusOK, user)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {

	id := request.Header.Get("id")
	idUser, _ := strconv.Atoi(id)
	user := db.Delete(idUser)

	responseWithJSON(response, http.StatusOK, user)

}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
