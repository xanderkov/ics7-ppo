package service

import (
	"context"
	"hospital/internal/modules/domain/dayPlan/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IDayPlanRepo

type IDayPlanRepo interface {
	GetById(ctx context.Context, id int32) (*dto.DayPlan, error)
	List(ctx context.Context) (dto.DayPlans, error)
	Create(ctx context.Context, dtm *dto.CalculateDay) (*dto.DayPlan, error)
	Update(ctx context.Context, num int32, dtm *dto.ChangeDay) (*dto.DayPlan, error)
	Delete(ctx context.Context, num int32) error
}

type DayPlanService struct {
	repo IDayPlanRepo
}

func NewDayPlanService(repo IDayPlanRepo) *DayPlanService {
	return &DayPlanService{
		repo: repo,
	}
}

func (r *DayPlanService) GetById(ctx context.Context, id int32) (*dto.DayPlan, error) {
	return r.repo.GetById(ctx, id)
}

func (r *DayPlanService) List(ctx context.Context) (dto.DayPlans, error) {
	return r.repo.List(ctx)
}

func (r *DayPlanService) Create(ctx context.Context, dtm *dto.CalculateDay) (*dto.DayPlan, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *DayPlanService) Update(ctx context.Context, id int32, dtm *dto.ChangeDay) (*dto.DayPlan, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *DayPlanService) Delete(ctx context.Context, id int32) error {
	return r.repo.Delete(ctx, id)
}
