package facade

import (
	"coincap/internal/entity"
	"coincap/internal/usecase"
	"context"
)

type (
	UserFacadeItf interface {
		SignUp(ctx context.Context, req entity.User) error
		Login(ctx context.Context, req entity.LoginRequest) (string, error)
		Logout(ctx context.Context, req entity.LogoutRequest) error
	}

	userFacade struct {
		userUC usecase.UserUsecaseItf
	}

	UserFacadeParams struct {
		UserUC usecase.UserUsecaseItf
	}
)

func NewUserFacade(params UserFacadeParams) UserFacadeItf {
	return &userFacade{
		userUC: params.UserUC,
	}
}

func (u *userFacade) SignUp(ctx context.Context, req entity.User) error {
	return u.userUC.SignUp(ctx, req)
}

func (u *userFacade) Login(ctx context.Context, req entity.LoginRequest) (string, error) {
	return u.userUC.Login(ctx, req)
}

func (u *userFacade) Logout(ctx context.Context, req entity.LogoutRequest) error {
	return u.userUC.Logout(ctx, req)
}
