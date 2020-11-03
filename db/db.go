package db

import (
	"fmt"

	"database/sql"

	"github.com/PhongVX/golang-rest-api/entities"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pass   = "991999"
	dbname = "PTUDW"
)

func GetData(username, password string) entities.User {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var user entities.User
	query := fmt.Sprintf("select username, email, phone from \"Users\" where username = '%s';", username)
	err = db.QueryRow(query).Scan(&user.Username, &user.Email, &user.Phone)
	return user
}
