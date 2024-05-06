package facade

import (
	"coincap/internal/constant"
	"coincap/internal/entity"
	"coincap/internal/usecase"
	"coincap/pkg/converter"
	"context"
	"sync"
)

type (
	CoinFacadeItf interface {
		GetCoins(ctx context.Context) ([]entity.Coin, error)
	}

	coinFacade struct {
		coinCapUC usecase.CoinCapUsecaseItf
	}

	CoinFacadeParams struct {
		CoinCapUC usecase.CoinCapUsecaseItf
	}
)

func NewCoinFacade(params CoinFacadeParams) CoinFacadeItf {
	return &coinFacade{
		coinCapUC: params.CoinCapUC,
	}
}

func (c *coinFacade) GetCoins(ctx context.Context) ([]entity.Coin, error) {
	rates, err := c.coinCapUC.GetRateByID(ctx, constant.CoinCapRatesIDR)
	if err != nil {
		return nil, err
	}

	idrRates := constant.USDRate / converter.ToFloat64(rates.RateUSD)

	listOfCoins, err := c.coinCapUC.GetAssets(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]entity.Coin, len(listOfCoins))
	wg := &sync.WaitGroup{}
	for i := range listOfCoins {
		wg.Add(1)
		go func(idx int, goWG *sync.WaitGroup) {
			defer goWG.Done()

			res[idx] = convertAssetToCoin(listOfCoins[idx], idrRates)
		}(i, wg)
	}
	wg.Wait()

	return res, nil
}
