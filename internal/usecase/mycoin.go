package usecase

import (
	"coincap/internal/entity"
	"coincap/internal/repository"
	"context"
)

type (
	MyCoinUsecaseItf interface {
		AddStaredCoin(ctx context.Context, req entity.StaredCoin) error
		GetStaredCoin(ctx context.Context, userID int64) (entity.StaredCoin, error)
		UpdateStaredCoin(ctx context.Context, req entity.StaredCoin) error
	}

	myCoinUseCase struct {
		myCoinRepository repository.MyCoinRepositoryItf
	}

	MyCoinUseCaseParams struct {
		MyCoinRepository repository.MyCoinRepositoryItf
	}
)

func NewMyCoinUsecase(params MyCoinUseCaseParams) MyCoinUsecaseItf {
	return &myCoinUseCase{
		myCoinRepository: params.MyCoinRepository,
	}
}

func (m *myCoinUseCase) AddStaredCoin(ctx context.Context, req entity.StaredCoin) error {
	return m.myCoinRepository.AddStaredCoin(ctx, req)
}

func (m *myCoinUseCase) GetStaredCoin(ctx context.Context, userID int64) (entity.StaredCoin, error) {
	return m.myCoinRepository.GetStaredCoin(ctx, userID)
}

func (m *myCoinUseCase) UpdateStaredCoin(ctx context.Context, req entity.StaredCoin) error {
	return m.myCoinRepository.UpdateStaredCoin(ctx, req)
}
