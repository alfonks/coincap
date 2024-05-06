package usecase

import (
	"coincap/internal/entity"
	"coincap/internal/gateway"
	"context"
)

type (
	CoinCapUsecaseItf interface {
		GetAssets(ctx context.Context) ([]entity.CoinCapAsset, error)
		GetAssetByID(ctx context.Context, id string) (entity.CoinCapAsset, error)
		GetRateByID(ctx context.Context, id string) (entity.CoinCapRates, error)
	}

	coinCapUseCase struct {
		coinCapGateway gateway.CoinCapGatewayItf
	}

	CoinCapUseCaseParams struct {
		CoinCapGateway gateway.CoinCapGatewayItf
	}
)

func NewCoinUsecase(params CoinCapUseCaseParams) CoinCapUsecaseItf {
	return &coinCapUseCase{
		coinCapGateway: params.CoinCapGateway,
	}
}

func (c *coinCapUseCase) GetAssets(ctx context.Context) ([]entity.CoinCapAsset, error) {
	return c.coinCapGateway.GetAssets(ctx)
}

func (c *coinCapUseCase) GetAssetByID(ctx context.Context, id string) (entity.CoinCapAsset, error) {
	return c.coinCapGateway.GetAssetByID(ctx, id)
}

func (c *coinCapUseCase) GetRateByID(ctx context.Context, id string) (entity.CoinCapRates, error) {
	return c.coinCapGateway.GetRateByID(ctx, id)
}
