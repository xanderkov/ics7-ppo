package modules

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"hospital/src/internal/modules/app"
	"hospital/src/internal/modules/config"
	"hospital/src/internal/modules/db"
	"hospital/src/internal/modules/domain"
	"hospital/src/internal/modules/logger"
)

var (
	AppModule = fx.Options(
		app.Module,
		logger.Module,
		config.Module,
		db.Module,

		domain.Module,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	AppInvokables = fx.Options(
		app.Invokables,
		logger.Invokables,
		config.Invokables,
		db.Invokables,

		domain.Invokables,
	)
)
