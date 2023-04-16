package domain

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/auth"
	"hospital/internal/modules/domain/doctor"
	"hospital/internal/modules/domain/patient"
	"hospital/internal/modules/domain/room"
)

var (
	Module = fx.Options(
		doctor.Module,
		auth.Module,
		room.Module,
		patient.Module,
	)
	Invokables = fx.Options(

		doctor.Invokables,
		auth.Invokables,
		room.Invokables,
		patient.Module,
	)
)
