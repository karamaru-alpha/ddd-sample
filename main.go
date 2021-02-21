package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

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
