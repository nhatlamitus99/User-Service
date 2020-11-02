package userapi

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/PhongVX/golang-rest-api/auth"
	"github.com/PhongVX/golang-rest-api/entities"
)

func Authorize(response http.ResponseWriter, request *http.Request) {

	var token_request entities.TokenRequest
	err := json.NewDecoder(request.Body).Decode(&token_request)

	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	} else {

		if token_request.Grant_Type != "password" {
			responseWithError(response, http.StatusForbidden, "forbidden password grant")
		} else {
			token, err := auth.CreateToken(token_request.Username, token_request.Password)
			if err != nil {
				responseWithError(response, http.StatusForbidden, err.Error())
			} else {
				requestBody, err := json.Marshal(map[string]string{
					"access_token": token,
					"token_type":   "Bearer",
					"expires_in":   "3600",
				})

				if err != nil {
					responseWithError(response, http.StatusBadRequest, err.Error())
				} else {
					resp, err := http.Post("http://127.0.0.1:4000/api/resource", "application/json", bytes.NewBuffer(requestBody))
					if err != nil {
						responseWithError(response, http.StatusBadRequest, err.Error())
					}
					defer resp.Body.Close()

				}
			}
		}

	}
	// pgdb := db.GetDB()

	// var user entities.User
	// data := pgdb.Select("username").Find(&user)

	// fmt.Println(data)

}

func GetResource(response http.ResponseWriter, request *http.Request) {
	err := auth.TokenValid(request)
	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	} else {
		responseWithJSON(response, http.StatusOK, "get protected resource successfully")
	}
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
