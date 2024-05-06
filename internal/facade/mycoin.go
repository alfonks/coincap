package facade

import (
	"coincap/internal/constant"
	"coincap/internal/entity"
	"coincap/internal/usecase"
	"coincap/pkg/converter"
	"context"
	"fmt"
	"log"
)

type (
	MyCoinFacadeItf interface {
		GetStaredCoin(ctx context.Context, userID int64) ([]entity.Coin, error)
		AddStaredCoin(ctx context.Context, req entity.AddStaredCoinRequest) error
		DeleteStaredCoin(ctx context.Context, req entity.DeleteStaredCoinRequest) error
	}

	myCoinFacade struct {
		myCoinUC  usecase.MyCoinUsecaseItf
		userUC    usecase.UserUsecaseItf
		coinCapUC usecase.CoinCapUsecaseItf
	}

	MyCoinFacadeParams struct {
		MyCoinUC  usecase.MyCoinUsecaseItf
		UserUC    usecase.UserUsecaseItf
		CoinCapUC usecase.CoinCapUsecaseItf
	}
)

func NewMyCoinFacade(params MyCoinFacadeParams) MyCoinFacadeItf {
	return &myCoinFacade{
		myCoinUC:  params.MyCoinUC,
		userUC:    params.UserUC,
		coinCapUC: params.CoinCapUC,
	}
}

func (m *myCoinFacade) GetStaredCoin(ctx context.Context, userID int64) ([]entity.Coin, error) {
	funcName := "facade.(*myCoinFacade).GetStaredCoin"
	userData, err := m.userUC.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if userData.ID == 0 {
		return nil, fmt.Errorf("[%v] user didn't exist, please create an account first", funcName)
	}

	rates, err := m.coinCapUC.GetRateByID(ctx, constant.CoinCapRatesIDR)
	if err != nil {
		return nil, err
	}

	idrRates := constant.USDRate / converter.ToFloat64(rates.RateUSD)

	staredCoin, err := m.myCoinUC.GetStaredCoin(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]entity.Coin, 0, len(staredCoin.StaredCoin))

	for assetID := range staredCoin.StaredCoin {
		asset, err := m.coinCapUC.GetAssetByID(ctx, assetID)
		if err != nil {
			log.Printf("[%v] error get asset for id: %v", funcName, assetID)
			asset.ID = assetID
		}
		res = append(res, convertAssetToCoin(asset, idrRates))
	}

	return res, nil
}

func (m *myCoinFacade) AddStaredCoin(ctx context.Context, req entity.AddStaredCoinRequest) error {
	funcName := "facade.(*myCoinFacade).AddStaredCoin"

	userData, err := m.userUC.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if userData.ID == 0 {
		return fmt.Errorf("[%v] user didn't exist, please create an account first", funcName)
	}

	staredCoin, err := m.myCoinUC.GetStaredCoin(ctx, req.UserID)
	if err != nil {
		return err
	}

	if staredCoin.ID == 0 {
		staredCoin.StaredCoin = make(map[string]bool)
		staredCoin.StaredCoin[req.CoinID] = true
		staredCoin.UserID = req.UserID

		err = m.myCoinUC.AddStaredCoin(ctx, staredCoin)
		if err != nil {
			return err
		}
	}

	staredCoin.StaredCoin[req.CoinID] = true
	err = m.myCoinUC.UpdateStaredCoin(ctx, staredCoin)
	if err != nil {
		return err
	}

	return nil
}

func (m *myCoinFacade) DeleteStaredCoin(ctx context.Context, req entity.DeleteStaredCoinRequest) error {
	funcName := "facade.(*myCoinFacade).DeleteStaredCoin"

	userData, err := m.userUC.GetUserByID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if userData.ID == 0 {
		return fmt.Errorf("[%v] user didn't exist, please create an account first", funcName)
	}

	staredCoin, err := m.myCoinUC.GetStaredCoin(ctx, req.UserID)
	if err != nil {
		return err
	}

	if staredCoin.ID == 0 {
		return fmt.Errorf("[%v] stared coin didn't exist, please add a new stared coin first", funcName)
	}

	delete(staredCoin.StaredCoin, req.CoinID)
	err = m.myCoinUC.UpdateStaredCoin(ctx, staredCoin)
	if err != nil {
		return err
	}

	return nil
}
