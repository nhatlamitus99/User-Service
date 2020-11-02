package userapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/PhongVX/golang-rest-api/auth"
	"github.com/PhongVX/golang-rest-api/db"
	"github.com/PhongVX/golang-rest-api/entities"
)

func Authorize(response http.ResponseWriter, request *http.Request) {

	var token_request entities.TokenRequest
	err := json.NewDecoder(request.Body).Decode(&token_request)

	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	} else {

		if token_request.Grant_Type != "password" || !checkPermit(token_request.Username, token_request.Password) {
			responseWithError(response, http.StatusForbidden, "forbidden password grant")
		} else {
			token, err := auth.CreateToken(token_request.Username, token_request.Password)
			if err != nil {
				responseWithError(response, http.StatusForbidden, err.Error())
			} else {
				requestBody, err := json.Marshal(map[string]string{
					"access_token": token,
					"token_type":   "Bearer",
					"expires_in":   string(time.Now().Add(time.Hour * 1).Unix()),
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

}

func GetResource(response http.ResponseWriter, request *http.Request) {
	err := auth.TokenValid(request)
	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	} else {
		data := db.GetDB("nhatlam", "")
		responseWithJSON(response, http.StatusOK, data)
	}
}

func checkPermit(username, password string) bool {
	user := db.GetDB(username, password)
	if user.Username == "" {
		return false
	}
	return true

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
