package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	infra2 "github.com/xxarupakaxx/linklist/infra"
	"os"
)

func init() {
	if os.Getenv("ISPRODUCTION") =="" {
		err := godotenv.Load(".env")
		if err != nil {
			logrus.Fatalln("error .env loading",err)
		}
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"https://linklistliff.herokuapp.com"}}))

	db,err := infra2.DBConnect()
	if err != nil {
		logrus.Infof("error connecting DB: %w",err)

		db, _ = infra2.DBConnect()
	}
	dbc, err := db.DB()
	defer func(dbc *sql.DB) {
		err := dbc.Close()
		if err != nil {
			logrus.Fatalf("error close db :%w",err)
		}
	}(dbc)

	db.Logger.LogMode()


}
