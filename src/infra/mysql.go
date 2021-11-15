package infra

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func DBConnect()(db *gorm.DB, err error) {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHostname := os.Getenv("DB_HOSTNAME")
	dbOption := "?parseTime=true&loc=Asia/Tokyo"
	dsn := dbUser+":"+dbPass+"@tcp("+dbHostname+":3306)/"+dbName+dbOption
	db,err = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed in Connect DB:%w", err)
	}
	return db,nil
}
