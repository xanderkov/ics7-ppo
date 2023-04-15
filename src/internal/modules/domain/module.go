package domain

import (
	"go.uber.org/fx"
	"hospital/src/internal/modules/domain/doctor/repo"
	"hospital/src/internal/modules/domain/doctor/service"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.DoctorRepo) *repo.DoctorRepo { return r },
				fx.As(new(service.IDoctorRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
