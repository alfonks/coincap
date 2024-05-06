package usecase

import (
	"coincap/internal/constant"
	"coincap/internal/entity"
	"coincap/internal/repository"
	"coincap/pkg/cfg"
	"coincap/pkg/converter"
	"coincap/pkg/encrypt"
	"coincap/pkg/random"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type (
	UserUsecaseItf interface {
		SignUp(ctx context.Context, req entity.User) error
		GetUserByID(ctx context.Context, id int64) (entity.User, error)
		Login(ctx context.Context, req entity.LoginRequest) (string, error)
		Logout(ctx context.Context, req entity.LogoutRequest) error
	}

	userUseCase struct {
		userRepo repository.UserRepositoryItf
		config   *cfg.ConfigSchema
	}

	UserUseCaseParams struct {
		UserRepo repository.UserRepositoryItf
		Config   *cfg.ConfigSchema
	}
)

func NewUserUseCase(params UserUseCaseParams) UserUsecaseItf {
	return &userUseCase{
		userRepo: params.UserRepo,
		config:   params.Config,
	}
}

func (u *userUseCase) SignUp(ctx context.Context, req entity.User) error {
	req.Password = encrypt.HashValue(req.Password, u.config.Secret.HashKey)

	userData, err := u.userRepo.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if userData.Email == req.Email {
		return errors.New("User email already exist")
	}

	if userData.Username == req.Username {
		return errors.New("User username already exist")
	}

	return u.userRepo.SignUp(ctx, req)
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (entity.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}

func (u *userUseCase) Login(ctx context.Context, req entity.LoginRequest) (string, error) {
	funcName := "usecase.(*userUseCase).Login"

	userData, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}

	if userData.ID == 0 {
		log.Printf("[%v] no data found for specified user: %v\n", funcName, req.Username)
		return "", fmt.Errorf("[%v] invalid username or password", funcName)
	}

	pw := encrypt.HashValue(req.Password, u.config.Secret.HashKey)

	if userData.Password != pw {
		log.Printf("[%v] wrong password for user: %v\n", funcName, req.Username)
		return "", fmt.Errorf("[%v] invalid username or password", funcName)
	}

	jwtToken := jwt.New(jwt.SigningMethodHS512)
	claims := jwtToken.Claims.(jwt.MapClaims)
	seed := random.RandomString()
	claims[constant.KeyTokenUserID] = userData.ID
	claims[constant.KeyTokenEmail] = userData.Email
	claims[constant.KeyTokenUsername] = userData.Username
	claims[constant.KeyTokenExpiryDate] = converter.ToString((time.Now().Add(time.Duration(u.config.JWT.ExpiryTime) * time.Hour)).Unix())
	claims[constant.KeyTokenRandomSeed] = seed
	jwtToken.Claims = claims

	token, err := jwtToken.SignedString([]byte(u.config.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	err = u.userRepo.UpdateLoginData(ctx, entity.User{
		ID:        userData.ID,
		IsLogedIn: true,
		Seed:      seed,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userUseCase) Logout(ctx context.Context, req entity.LogoutRequest) error {
	funcName := "usecase.(*userUseCase).Logout"
	userData, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return err
	}

	if userData.ID == 0 {
		log.Printf("[%v] no data found for specified user: %v\n", funcName, req.Username)
		return fmt.Errorf("[%v] invalid username or password", funcName)
	}

	err = u.userRepo.UpdateLoginData(ctx, entity.User{
		ID: userData.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
