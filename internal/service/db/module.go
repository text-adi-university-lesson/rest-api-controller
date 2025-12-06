package db

import (
	"go.uber.org/fx"
)

// Module exports the DB provider to the app's Fx graph.
var Module = fx.Module("db",
	fx.Provide(
		NewConfig,
		NewGormDB,
	),
)
