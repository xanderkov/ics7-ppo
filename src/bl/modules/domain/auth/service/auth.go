package service

import (
	"context"
	"hospital/src/bl/modules/dayPlan/dto"
	"hospital/src/bl/modules/doctor/dto"
)


type IUserRepo interface {
	GetByTokenId(ctx context.Context, tokenId int32) (*doctor_dto.Doctor, error)
	Create(ctx context.Context, dtm *doctor_dto.CreateDoctor) (*doctor_dto.Doctor, error)
}

type AuthService struct {
	repo   IUserRepo
	tokenId int32
}

func (r *AuthService) SignUp(ctx context.Context, newUser *dto.NewUser) (*doctor_dto.Doctor, error) {

	createUser := &doctor_dto.CreateUser{
		Suranme:      newUser.surname,
		TokenId:      newUser.tokenId,
		speciality    newUser.speciality,
		role          newUser.role
	}

	// Создаем пользователя
	createdUser, err := r.repo.Create(ctx, createUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}