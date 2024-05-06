package repository

import (
	"coincap/internal/entity"
	"coincap/pkg/converter"
	"coincap/pkg/database"
	"context"
	"encoding/json"
	"fmt"
)

type (
	MyCoinRepositoryItf interface {
		AddStaredCoin(ctx context.Context, req entity.StaredCoin) error
		GetStaredCoin(ctx context.Context, userID int64) (entity.StaredCoin, error)
		UpdateStaredCoin(ctx context.Context, req entity.StaredCoin) error
	}

	myCoinRepository struct {
		db database.DBItf
	}

	MyCoinRepositoryParams struct {
		DB database.DBItf
	}
)

func NewMyCoinRepository(params MyCoinRepositoryParams) MyCoinRepositoryItf {
	return &myCoinRepository{
		db: params.DB,
	}
}

func (m *myCoinRepository) AddStaredCoin(ctx context.Context, req entity.StaredCoin) error {
	funcName := "repository.(*UserRepository).AddStaredCoin"

	data, err := json.Marshal(req.StaredCoin)
	if err != nil {
		return fmt.Errorf("[%v] fail marshal stared coin for user: %v, err: %v", funcName, req.UserID, err)
	}

	req.StaredCoinStr = converter.ToString(data)

	conn := m.db
	query := `
		INSERT INTO user_stared_coin(
			user_id, 
			stared_coin 
		) VALUES (?, ?)
	`

	err = conn.Exec(
		query,
		req.UserID,
		req.StaredCoinStr,
	).Error()
	if err != nil {
		return fmt.Errorf("[%v] error insert stared coin for user: %v, error: %v", funcName, req.UserID, err)
	}

	return nil
}

func (m *myCoinRepository) GetStaredCoin(ctx context.Context, userID int64) (entity.StaredCoin, error) {
	funcName := "repository.(*UserRepository).GetStaredCoin"

	conn := m.db
	query := `
		SELECT 
			id, user_id, stared_coin
		FROM 
			user_stared_coin
		WHERE
			user_id = ?
	`

	var res entity.StaredCoin
	err := conn.Raw(
		query,
		userID,
	).Scan(&res).Error()
	if err != nil {
		return entity.StaredCoin{}, fmt.Errorf("[%v] error get stared coin by userID: %v, error: %v", funcName, userID, err)
	}

	if res.ID != 0 {
		staredCoin := make(map[string]bool)
		err = json.Unmarshal([]byte(res.StaredCoinStr), &staredCoin)
		if err != nil {
			return entity.StaredCoin{}, fmt.Errorf("[%v] error unmarshal stared coin by userID: %v, error: %v", funcName, userID, err)
		}
		res.StaredCoin = staredCoin
		res.StaredCoinStr = ""
	}

	return res, nil
}

func (m *myCoinRepository) UpdateStaredCoin(ctx context.Context, req entity.StaredCoin) error {
	funcName := "repository.(*UserRepository).UpdateStaredCoin"

	data, err := json.Marshal(req.StaredCoin)
	if err != nil {
		return fmt.Errorf("[%v] fail marshal stared coin for user: %v, err: %v", funcName, req.UserID, err)
	}

	req.StaredCoinStr = converter.ToString(data)

	conn := m.db
	query := `
		UPDATE 
			user_stared_coin 
		SET
			stared_coin = ?
		WHERE
			user_id = ?
	`

	err = conn.Exec(
		query,
		req.StaredCoinStr,
		req.UserID,
	).Error()
	if err != nil {
		return fmt.Errorf("[%v] error update stared coin for user: %v, error: %v", funcName, req.UserID, err)
	}

	return nil
}
