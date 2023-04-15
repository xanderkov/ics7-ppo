package main

import (
	"go.uber.org/fx"
	"hospital/src/internal/modules"
)

func main() {
	fx.New(
		modules.AppModule,
		modules.AppInvokables,
	).Run()
}
