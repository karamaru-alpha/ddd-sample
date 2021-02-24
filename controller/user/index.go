package user

import (
	"net/http"

	"github.com/labstack/echo"

	applicationService "github.com/karamaru-alpha/ddd-sample/application/user"
)

// IController Interface of Controller handle request
type IController interface {
	Create(echo.Context) error
}

type controller struct {
	applicationService applicationService.IApplicationService
}

// NewController Fuction create user controller
func NewController(userApplicationService applicationService.IApplicationService) IController {
	return &controller{
		applicationService: userApplicationService,
	}
}

// Create Controller create user
func (uc controller) Create(c echo.Context) error {
	type (
		request struct {
			Name       string `json:"name"`
			MailAdress string `json:"mail_adress"`
		}
		response struct {
			AuthToken string `json:"auth_token"`
		}
	)

	requestBody := new(request)
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	if err := uc.applicationService.Register(requestBody.Name, requestBody.MailAdress); err != nil {
		return err
	}

	// TODO implement JWT
	return c.JSON(http.StatusCreated, &response{
		AuthToken: "foo bar token",
	})
}
