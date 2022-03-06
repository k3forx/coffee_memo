package user

import "context"

func NewUsecase() *UserUsecase {
	return &UserUsecase{}
}

type Usecase interface {
}

type UserUsecase struct{}

var _ Usecase = (*Usecase)(nil)

func (u *UserUsecase) GetByID(ctx context.Context) {}
