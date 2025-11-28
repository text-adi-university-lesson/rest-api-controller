package http

import (
	"rest-api-controller/internal/module/sensors"
	"rest-api-controller/internal/service/fiber"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func Run(cmd *cobra.Command) {
	fx.New(
		fx.Provide(func() *cobra.Command { return cmd }),

		sensors.Module,

		fiber.Module,
	).Run()
}
