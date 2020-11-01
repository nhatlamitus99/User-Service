package userapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PhongVX/golang-rest-api/models"

	"Client/db"

	"github.com/PhongVX/golang-rest-api/entities"
)

type User struct {
	password string
	username string
}

func Authorize(response http.ResponseWriter, request *http.Request) {

	request.ParseForm()

	var user User
	for key, value := range request.Form {
		if key == "username" {
			user.username = value[0]
		} else {
			user.password = value[0]
		}
	}

	log.Println(user.username, user.password)

	pgdb := db.GetDB()
	_, err := pgdb.Query("SELECT * FROM USER")

	log.Println(err)

	if user.password == "123" && user.username == "nhatlam" {
		responseWithJSON(response, http.StatusOK, user.username)
	} else {
		responseWithJSON(response, http.StatusForbidden, "failed")
	}

}

func GetInfo(response http.ResponseWriter, request *http.Request) {
	users := models.GetAllUser()
	responseWithJSON(response, http.StatusOK, users)
}

func RedirectToAuthServer(response http.ResponseWriter, request *http.Request) {

	redirect_uri := "http://localhost:3000/authorize/callback"
	response_type := "code"
	client_id := "infoCustomer"

	var path string
	path = "http://localhost:3001/oauth/authorize?redirect_uri=" + redirect_uri + "?response_type=" + response_type + "?client_id=" + client_id
	fmt.Println(path)

	http.Redirect(response, request, path, 301)
}

func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param 'id' is missing")
		return
	}
	user, err := models.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Couldn't create user")
			return
		}
		responseWithJSON(response, http.StatusOK, user)
	}
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.UpdateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Couldn't update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param 'id' is missing")
		return
	}
	result := models.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Couldn't delete user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete user successfully")
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
