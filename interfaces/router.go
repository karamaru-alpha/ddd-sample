package interfaces

import (
	"github.com/labstack/echo"

	userController "github.com/karamaru-alpha/ddd-sample/interfaces/controller/user"
)

// InitializeRoutes Define routes
func InitializeRoutes(e *echo.Echo, userController userController.IController) {
	e.POST("/register", userController.Create)
}
