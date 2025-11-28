package sensors

import (
	"rest-api-controller/internal/module/sensors/repository"
	"rest-api-controller/internal/module/sensors/service"
	httpController "rest-api-controller/internal/service/fiber/controller"

	"go.uber.org/fx"
)

// Params for DI
type Params struct {
	fx.In

	Service *service.Service
}

// Result for DI group
type Result struct {
	fx.Out

	Controller httpController.Controller `group:"controller"`
}

func New(p Params) Result {
	return Result{
		Controller: NewController(*p.Service),
	}
}

var Module = fx.Module("entity",
	// Controller -> Service -> Repository
	fx.Provide(
		New,
	),
	fx.Provide(
		service.NewService,
	),
	fx.Provide(
		repository.NewLocalRepository,
	),
)
