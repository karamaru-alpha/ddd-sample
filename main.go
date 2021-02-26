package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/karamaru-alpha/ddd-sample/infrastructure/mysql"

	userApplicationService "github.com/karamaru-alpha/ddd-sample/application/user"
	userController "github.com/karamaru-alpha/ddd-sample/controller/user"
	userDomainService "github.com/karamaru-alpha/ddd-sample/domain/service/user"
	userFactory "github.com/karamaru-alpha/ddd-sample/factory/user"
	userRepositoryImpl "github.com/karamaru-alpha/ddd-sample/infrastructure/mysql/repository_impl/user"
)

func main() {
	// TODO implement DI by google/wire
	userRepositoryImpl := userRepositoryImpl.NewRepositoryImpl(mysql.Conn)
	userDomainService := userDomainService.NewDomainService(userRepositoryImpl)
	userFactory := userFactory.NewFactory()
	userApplicationService := userApplicationService.NewApplicationService(userDomainService, userRepositoryImpl, userFactory)
	userController := userController.NewController(userApplicationService)

	// TODO Implement server init function on another layer.
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.POST("/register", userController.Create)

	log.Println("Server running...")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
