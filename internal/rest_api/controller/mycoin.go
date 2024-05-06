package controller

import (
	"coincap/internal/entity"
	"coincap/internal/facade"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyCoinCtrl struct {
	myCoinFacade facade.MyCoinFacadeItf
}

type MyCoinCtrlParams struct {
	MyCoinFacade facade.MyCoinFacadeItf
}

func NewMyCoinCtrl(params MyCoinCtrlParams) *MyCoinCtrl {
	return &MyCoinCtrl{
		myCoinFacade: params.MyCoinFacade,
	}
}

func (m *MyCoinCtrl) GetStaredCoin(c *gin.Context) {
	funcName := "controller.(*MyCoinCtrl).AddStaredCoin"
	ctx := c.Request.Context()

	var req entity.GetStaredCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[%v] error bind request, err: %v", funcName, err)
		c.JSON(http.StatusBadRequest, entity.GetStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	err := req.Validate()
	if err != nil {
		log.Printf("[%v] error validate request, err: %v", funcName, err)
		c.JSON(http.StatusBadRequest, entity.GetStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := m.myCoinFacade.GetStaredCoin(ctx, req.UserID)
	if err != nil {
		log.Printf("[%v] error get stared coin for user: %v, err: %v", funcName, req.UserID, err)
		c.JSON(http.StatusInternalServerError, entity.GetStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.GetStaredCoinResponse{
		Message: "success",
		Data:    res,
	})
}

func (m *MyCoinCtrl) AddStaredCoin(c *gin.Context) {
	funcName := "controller.(*MyCoinCtrl).AddStaredCoin"
	ctx := c.Request.Context()

	var req entity.AddStaredCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[%v] error bind request, err: %v", funcName, err)
		c.JSON(http.StatusBadRequest, entity.AddStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	err := req.Validate()
	if err != nil {
		log.Printf("[%v] error validate request, err: %v", funcName, err)
		c.JSON(http.StatusBadRequest, entity.AddStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	err = m.myCoinFacade.AddStaredCoin(ctx, req)
	if err != nil {
		log.Printf("[%v] error add stared coin for user: %v, err: %v", funcName, req.UserID, err)
		c.JSON(http.StatusInternalServerError, entity.AddStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.AddStaredCoinResponse{
		Message: "success",
	})
}

func (m *MyCoinCtrl) DeleteStaredCoin(c *gin.Context) {
	funcName := "controller.(*MyCoinCtrl).AddStaredCoin"
	ctx := c.Request.Context()

	var req entity.DeleteStaredCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[%v] error bind request, err: %v", funcName, err)
		c.JSON(http.StatusBadRequest, entity.DeleteStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	err := req.Validate()
	if err != nil {
		log.Printf("[%v] error validate request, err: %v", funcName, err)
		c.JSON(http.StatusBadRequest, entity.DeleteStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	err = m.myCoinFacade.DeleteStaredCoin(ctx, req)
	if err != nil {
		log.Printf("[%v] error delete stared coin for user: %v, err: %v", funcName, req.UserID, err)
		c.JSON(http.StatusInternalServerError, entity.DeleteStaredCoinResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.DeleteStaredCoinResponse{
		Message: "success",
	})
}
