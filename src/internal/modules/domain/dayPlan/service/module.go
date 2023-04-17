package service

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewDayPlanService)
	Invokables = fx.Invoke()
)
