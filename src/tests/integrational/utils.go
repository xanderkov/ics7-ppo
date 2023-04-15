package integrational

import (
	"context"
	"github.com/google/uuid"
	"hospital/internal/models/session"
	"hospital/internal/modules/domain/doctor/dto"
)

func makeCtxByUser(doctor *dto.Doctor) context.Context {
	ss := session.Session{
		SessionID: uuid.NewString(),
		UserId:    doctor.Id,
	}

	ctx := context.Background()
	return session.SetSessionToCtx(ctx, ss)
}
