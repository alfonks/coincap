package entity

import "errors"

type (
	AddStaredCoinRequest struct {
		CoinID string `json:"coin_id"`
		UserID int64  `json:"user_id"`
	}

	AddStaredCoinResponse struct {
		Message string `json:"message"`
	}

	DeleteStaredCoinRequest struct {
		CoinID string `json:"coin_id"`
		UserID int64  `json:"user_id"`
	}

	DeleteStaredCoinResponse struct {
		Message string `json:"message"`
	}

	GetStaredCoinRequest struct {
		UserID int64 `json:"user_id"`
	}

	GetStaredCoinResponse struct {
		Message string `json:"message"`
		Data    []Coin `json:"data"`
	}

	StaredCoin struct {
		ID            int64           `json:"id"`
		UserID        int64           `json:"user_id"`
		StaredCoin    map[string]bool `json:"stared_coin" gorm:"-"`
		StaredCoinStr string          `json:"-" gorm:"column:stared_coin"`
	}
)

func (a *AddStaredCoinRequest) Validate() error {
	if a.CoinID == "" {
		return errors.New("No new coin to add")
	}

	if a.UserID == 0 {
		return errors.New("Missing user id")
	}

	return nil
}

func (d *DeleteStaredCoinRequest) Validate() error {
	if d.CoinID == "" {
		return errors.New("No new coin to add")
	}

	if d.UserID == 0 {
		return errors.New("Missing user id")
	}

	return nil
}

func (g *GetStaredCoinRequest) Validate() error {
	if g.UserID == 0 {
		return errors.New("Missing user id")
	}

	return nil
}
