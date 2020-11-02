package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "991999"
	dbname   = "PTUDW"
)

func GetDB() *gorm.DB {

	//dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	db, _ := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	return db
}
