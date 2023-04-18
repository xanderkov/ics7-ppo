package view

import (
	"go.uber.org/fx"
	"hospital/internal/modules/view/controllers"
)

var (
	Module     = fx.Provide(controllers.NewController)
	Invokables = fx.Invoke(startBot)
)
