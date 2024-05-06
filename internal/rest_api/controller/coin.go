package controller

import (
	"coincap/internal/entity"
	"coincap/internal/facade"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoinCtrl struct {
	coinFacade facade.CoinFacadeItf
}

type CoinCtrlParams struct {
	CoinFacade facade.CoinFacadeItf
}

func NewCoinCtrl(params CoinCtrlParams) *CoinCtrl {
	return &CoinCtrl{
		coinFacade: params.CoinFacade,
	}
}

func (coin *CoinCtrl) GetCoins(c *gin.Context) {
	funcName := "controller.(*CoinCtrl).GetCoins"
	ctx := c.Request.Context()

	coins, err := coin.coinFacade.GetCoins(ctx)
	if err != nil {
		log.Printf("[%v] error get coins, err: %v\n", funcName, err)
		c.JSON(http.StatusInternalServerError, entity.GetCoinsResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.GetCoinsResponse{
		Message: "success",
		Data:    coins,
	})
}
