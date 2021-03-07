package user

import (
	"net/http"

	"github.com/labstack/echo"

	applicationService "github.com/karamaru-alpha/ddd-sample/application/user/create"
)

// IController Interface of Controller handle request
type IController interface {
	Create(echo.Context) error
}

type controller struct {
	applicationService applicationService.IInputPort
}

// NewController Fuction create user controller
func NewController(applicationService applicationService.IInputPort) IController {
	return &controller{
		applicationService,
	}
}

// Create Controller create user
func (uc controller) Create(c echo.Context) error {
	type (
		request struct {
			Name        string `json:"name"`
			MailAddress string `json:"mail_address"`
			Password    string `json:"password"`
		}
		response struct {
			AuthToken string `json:"auth_token"`
		}
	)

	requestBody := new(request)
	if err := c.Bind(requestBody); err != nil {
		return err
	}

	inputData := applicationService.InputData{
		Name:          requestBody.Name,
		MailAddress:   requestBody.MailAddress,
		PlainPassword: requestBody.Password,
	}

	outputData := uc.applicationService.Handle(inputData)
	if outputData.Err != nil {
		return outputData.Err
	}

	return c.JSON(http.StatusCreated, &response{
		AuthToken: outputData.Token,
	})
}
