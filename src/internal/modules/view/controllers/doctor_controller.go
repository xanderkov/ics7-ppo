package controllers

import (
	"context"
	dto1 "hospital/internal/modules/domain/doctor/dto"
)

func (r *Controller) Doctor(ctx context.Context, id int) (*dto1.Doctor, error) {
	user, err := r.doctorService.GetById(ctx, id)

	return user, err
}
