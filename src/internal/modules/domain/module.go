package domain

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/auth"
	"hospital/internal/modules/domain/doctor"
)

var (
	Module = fx.Options(
		doctor.Module,
		auth.Module,
	)
	Invokables = fx.Options(

		doctor.Invokables,
		auth.Invokables,
	)
)
