package cors

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

	Middleware httpController.Controller `group:"middleware"`
}

func New(p Params) Result {
	return Result{
		Middleware: NewMiddleware(),
	}
}

var Module = fx.Module("CorsMiddleware",
	fx.Provide(
		New,
	),
)
