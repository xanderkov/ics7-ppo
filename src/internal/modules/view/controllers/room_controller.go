package controllers

import (
	"context"
	dto1 "hospital/internal/modules/domain/room/dto"
)

func (r *Controller) Room(ctx context.Context, id int) (*dto1.Room, error) {
	user, err := r.roomService.GetByNum(ctx, id)

	return user, err
}
