package service

import (
	"context"
	"hospital/src/internal/modules/domain/auth/dto"
	doctor_dto "hospital/src/internal/modules/domain/doctor/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IUserRepo

type IUserRepo interface {
	GetByTokenId(ctx context.Context, tokenId int32) (*doctor_dto.Doctor, error)
	Create(ctx context.Context, dtm *doctor_dto.CreateDoctor) (*doctor_dto.Doctor, error)
}

type AuthService struct {
	repo    IUserRepo
	tokenId int32
}

func (r *AuthService) SignUp(ctx context.Context, newUser *dto.NewUser) (*doctor_dto.Doctor, error) {

	createUser := &doctor_dto.CreateDoctor{
		Surname:    newUser.Surname,
		TokenId:    newUser.TokenId,
		Speciality: newUser.Speciality,
		Role:       newUser.Role,
	}

	// Создаем пользователя
	createdUser, err := r.repo.Create(ctx, createUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
