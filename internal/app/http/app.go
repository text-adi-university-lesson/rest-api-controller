package http

import (
	"rest-api-controller/internal/module/sensors"
	"rest-api-controller/internal/service/db"
	"rest-api-controller/internal/service/fiber"
	"rest-api-controller/internal/service/fiber/middleware/cors"
	"rest-api-controller/internal/service/fiber/middleware/healthcheck"
	"rest-api-controller/internal/service/fiber/middleware/recovery"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run(cmd *cobra.Command) {
	fx.New(
		fx.Provide(func() *cobra.Command { return cmd }),

		sensors.Module,

		recovery.Module,
		healthcheck.Module,
		cors.Module,

		db.Module,

		fiber.Module,
	).Run()
}
