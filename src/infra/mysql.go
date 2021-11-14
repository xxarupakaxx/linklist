package infra

import (
	"database/sql"
	"log"
	"os"
)

func DBConnect(db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbOption := "?parseTime=true&loc=Asia/Tokyo"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHostname+":3306)/"+dbName+dbOption)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err == nil {
		log.Println("1success")
	} else {
		log.Println("failed connect err:%w", err)
	}
}
