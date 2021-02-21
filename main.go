package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// TODO Implement infrastructure on another layer.
	db := sqlConnect()
	defer db.Close()

	// TODO Implement server init function on another layer.
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	log.Println("Server running...")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(ddd-sample-db:3306)"
	DBNAME := "ddd_sample"

	CONNECT := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=true&loc=Local",
		USER,
		PASS,
		PROTOCOL,
		DBNAME,
	)

	fmt.Print(CONNECT)

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}
