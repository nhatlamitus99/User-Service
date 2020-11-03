package userapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/PhongVX/golang-rest-api/auth"
	"github.com/PhongVX/golang-rest-api/db"
	"github.com/PhongVX/golang-rest-api/entities"
)

func Authorize(response http.ResponseWriter, request *http.Request) {

	var tokenRequest entities.TokenRequest
	err := json.NewDecoder(request.Body).Decode(&tokenRequest)

	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	} else {
		if tokenRequest.Grant_Type != "password" || !checkPermit(tokenRequest) {
			responseWithError(response, http.StatusForbidden, "Unauthorized")
		} else {
			os.Setenv("SECRET_KEY", tokenRequest.Client_Secret)
			token, err := auth.CreateToken(tokenRequest.Username, tokenRequest.Password)
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
	err, data := auth.TokenValid(request)
	if err != nil {
		responseWithError(response, http.StatusForbidden, err.Error())
	} else {
		resource := db.GetData(data.Username, data.Password)
		responseWithJSON(response, http.StatusOK, resource)
	}
}

func checkPermit(request entities.TokenRequest) bool {
	user := db.GetData(request.Username, request.Password)
	if user.Username != "nhatlam" {
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
