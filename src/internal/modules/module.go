package domain

import "go.uber.org/fx"

var (
	Module = fx.Options(
		auth.Module,
		auth.Module,
		event.Module,
		tag.Module,
		user.Module,
	)
	Invokables = fx.Options(
		access_right.Invokables,
		auth.Invokables,
		event.Invokables,
		tag.Invokables,
		user.Invokables,
	)
)
