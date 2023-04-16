package integrational

import (
	"context"

	"hospital/internal/modules/db/ent"
	doctor_servis "hospital/internal/modules/domain/doctor/service"
	patient_servis "hospital/internal/modules/domain/patient/service"
	room_servis "hospital/internal/modules/domain/room/service"

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
	doctorService *doctor_servis.DoctorService,
	// authService *auth_serv.AuthService,
	patientService *patient_servis.PatientService,
	roomService *room_servis.RoomService,

	client *ent.Client,
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				// doctorServiceTest(t, doctorService, authService, client)

				_ = shutdowner.Shutdown()
			}()

			return nil
		},
	})
}
