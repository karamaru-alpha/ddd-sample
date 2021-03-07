// +build wireinject

package user

import (
	"github.com/google/wire"

	"github.com/karamaru-alpha/ddd-sample/infrastructure/mysql"

	application "github.com/karamaru-alpha/ddd-sample/application/user/create"
	domainModel "github.com/karamaru-alpha/ddd-sample/domain/model/user"
	domainService "github.com/karamaru-alpha/ddd-sample/domain/service/user"
	repositoryImpl "github.com/karamaru-alpha/ddd-sample/infrastructure/mysql/user"
	controller "github.com/karamaru-alpha/ddd-sample/interfaces/controller/user"
)

// InitializeDI DI about user
func InitializeDI() controller.IController {
	wire.Build(
		controller.NewController,
		application.NewInteractor,
		repositoryImpl.NewRepositoryImpl,
		domainService.NewDomainService,
		domainModel.NewFactory,

		mysql.ConnectGorm,
	)

	return nil
}
