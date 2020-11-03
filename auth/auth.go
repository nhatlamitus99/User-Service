package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PhongVX/golang-rest-api/entities"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(username, password string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["password"] = password
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func TokenValid(request *http.Request) (error, entities.Data) {
	data := entities.Data{}
	tokenString := ExtractToken(request)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err, data
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		data.Username = claims["username"].(string)
		data.Password = claims["password"].(string)
	}
	return nil, data
}

func ExtractToken(request *http.Request) string {

	bearerToken := request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
