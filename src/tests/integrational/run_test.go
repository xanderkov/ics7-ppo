package integrational

import (
	"context"

	"hospital/internal/modules/db/ent"
	auth_serv "hospital/internal/modules/domain/auth/service"
	doctor_server "hospital/internal/modules/domain/doctor/service"

	"go.uber.org/fx"
	"testing"
)

func TestServices(t *testing.T) {
	fx.New(
		testModule,
		testInvokables,

		fx.Supply(t),
		fx.Invoke(execTests),
	).Run()

}

func execTests(
	t *testing.T,
	doctorService *doctor_server.DoctorService,
	authService *auth_serv.AuthService,

	client *ent.Client,
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				doctorServiceTest(t, doctorService, authService, client)

				_ = shutdowner.Shutdown()
			}()

			return nil
		},
	})
}
