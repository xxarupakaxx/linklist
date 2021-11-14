package infra

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"os"
)

func DBConnect(db *gorm.DB, err error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbOption := "?parseTime=true&loc=Asia/Tokyo"
	db, err = gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHostname+":3306)/"+dbName+dbOption)
	if err != nil {
		logrus.Fatalf("failed in Connect DB:%w", err)
	}
}
