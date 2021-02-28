package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	userDI "github.com/karamaru-alpha/ddd-sample/di/user"
	"github.com/karamaru-alpha/ddd-sample/interfaces"
)

func main() {
	// DI
	userController := userDI.InitializeDI()

	// Start Server
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	interfaces.InitializeRoutes(e, userController)

	log.Println("Server running...")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
