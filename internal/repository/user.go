package repository

import (
	"coincap/internal/entity"
	"coincap/pkg/database"
	"context"
	"fmt"
)

type (
	UserRepositoryItf interface {
		SignUp(ctx context.Context, req entity.User) error
		GetUserByEmailOrUsername(ctx context.Context, email, username string) (entity.User, error)
		GetUserByID(ctx context.Context, id int64) (entity.User, error)
		GetUserByUsername(ctx context.Context, username string) (entity.User, error)
		UpdateLoginData(ctx context.Context, req entity.User) error
	}

	userRepository struct {
		db database.DBItf
	}

	UserRepositoryParams struct {
		DB database.DBItf
	}
)

func NewUserRepository(params UserRepositoryParams) UserRepositoryItf {
	return &userRepository{
		db: params.DB,
	}
}

func (u *userRepository) SignUp(ctx context.Context, req entity.User) error {
	funcName := "repository.(*UserRepository).SignUp"

	conn := u.db
	query := `
		INSERT INTO user_data(
			name, 
			username, 
			email, 
			password
		) VALUES (?, ?, ?, ?)
	`

	err := conn.Exec(
		query,
		req.Name,
		req.Username,
		req.Email,
		req.Password,
	).Error()
	if err != nil {
		return fmt.Errorf("[%v] error creating new user, error: %v", funcName, err)
	}

	return nil
}

func (u *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (entity.User, error) {
	funcName := "repository.(*UserRepository).GetUserByEmailOrUsername"

	conn := u.db
	query := `
		SELECT 
			id, name, username, email, password, COALESCE(seed, '') as seed
		FROM 
			user_data
		WHERE
			email like ? OR username like ?
	`

	var res entity.User
	err := conn.Raw(
		query,
		email,
		username,
	).Scan(&res).Error()
	if err != nil {
		return entity.User{}, fmt.Errorf("[%v] error get user by email or username, error: %v", funcName, err)
	}

	return res, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, id int64) (entity.User, error) {
	funcName := "repository.(*UserRepository).GetUserByEmailOrUsername"

	conn := u.db
	query := `
		SELECT 
			id, name, username, email, password, COALESCE(seed, '') as seed
		FROM 
			user_data
		WHERE
			id = ?
	`

	var res entity.User
	err := conn.Raw(
		query,
		id,
	).Scan(&res).Error()
	if err != nil {
		return entity.User{}, fmt.Errorf("[%v] error get user by id, error: %v", funcName, err)
	}

	return res, nil
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	funcName := "repository.(*UserRepository).GetUserByUsername"

	conn := u.db
	query := `
		SELECT 
			id, name, username, email, password
		FROM 
			user_data
		WHERE
			username = ?
	`

	var res entity.User
	err := conn.Raw(
		query,
		username,
	).Scan(&res).Error()
	if err != nil {
		return entity.User{}, fmt.Errorf("[%v] error get user by username, error: %v", funcName, err)
	}

	return res, nil
}

func (u *userRepository) UpdateLoginData(ctx context.Context, req entity.User) error {
	funcName := "repository.(*UserRepository).UpdateStaredCoin"

	conn := u.db
	query := `
		UPDATE 
			user_data 
		SET
			is_loged_in = ?,
			seed = ?
		WHERE
			id = ?
	`

	err := conn.Exec(
		query,
		req.IsLogedIn,
		req.Seed,
		req.ID,
	).Error()
	if err != nil {
		return fmt.Errorf("[%v] error update login data user: %v, error: %v", funcName, req.ID, err)
	}

	return nil
}
