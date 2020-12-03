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

func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetUser(id int) entities.User {
	var user entities.User

	db, err := connectDB()
	if err != nil {
		return user
	}
	defer db.Close()

	query := fmt.Sprintf("select username, email, phone from \"Users\" where id = '%s';", id)
	err = db.QueryRow(query).Scan(&user.Username, &user.Email, &user.Phone)
	return user
}

func ListUser() []entities.User {
	// TODO
	return nil
}

func CreateUser(user entities.User) bool {
	// TODO
	return true
}

func UpdateUser(id int) bool {
	// TODO
	return true
}

func Delete(id int) bool {
	// TODO
	return true
}
