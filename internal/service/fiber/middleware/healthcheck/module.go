package healthcheck

import (
	httpController "rest-api-controller/internal/service/fiber/controller"

	"go.uber.org/fx"
)

type Params struct {
	fx.In
}

// Result for DI group
type Result struct {
	fx.Out

	Controller httpController.Controller `group:"controller"`
}

func New(p Params) Result {
	return Result{
		Controller: NewController(),
	}
}

var Module = fx.Module("CorsMiddleware",
	fx.Provide(
		New,
	),
)
