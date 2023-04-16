package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hospital/internal/modules/db/ent"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	auth_serv "hospital/internal/modules/domain/auth/service"
	"hospital/internal/modules/domain/patient/dto"
	"hospital/internal/modules/domain/patient/service"
	"testing"
)

func patientServiceTest(t *testing.T, service *service.PatientService, authService *auth_serv.AuthService, client *ent.Client) {

	// Регистрируем пользователя
	newUser := &auth_dto.NewDoctor{
		TokenId:    "1",
		Surname:    "Kovel",
		Speciality: "Психотерапевт",
		Role:       "Глав врач",
	}
	currentUser, err := authService.SignUp(context.Background(), newUser)
	assert.NoError(t, err)
	// Возврат, т.к. не получится создать сессию с пользователем
	if err != nil {
		return
	}

	// Создаем контекст с сессией
	ctx := makeCtxByUser(currentUser)

	newpatient := &dto.CreatePatient{
		Name:           "Alexander",
		Surname:        "Kovel",
		Patronymic:     "Denisovich",
		Height:         183,
		Weight:         80,
		RoomNumber:     1,
		DegreeOfDanger: 10,
	}
	patient, err := service.Create(ctx, newpatient)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	t1, err := service.GetById(ctx, patient.Id)
	assert.NoError(t, err)
	assert.Equal(t, patient, t1)

	ts1, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, patient, ts1[0])

	patient.Id += 1
	updateUser := &dto.UpdatePatient{
		Name:           patient.Name,
		Surname:        patient.Surname,
		Patronymic:     patient.Patronymic,
		Height:         patient.Height,
		Weight:         patient.Weight,
		RoomNumber:     patient.RoomNumber,
		DegreeOfDanger: patient.DegreeOfDanger,
	}
	t2, err := service.Update(ctx, patient.Id, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, patient, t2)

	err = service.Delete(ctx, patient.Id)
	assert.NoError(t, err)

	t3, err := service.GetById(ctx, patient.Id)
	assert.NoError(t, err)
	assert.Equal(t, patient, t3)

	ts2, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)

	ts3, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, patient, ts3[0])
}
