package controller

import (
	"coincap/internal/entity"
	"coincap/internal/facade"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserCtrl struct {
		userFacade facade.UserFacadeItf
	}

	UserCtrlParams struct {
		UserFacade facade.UserFacadeItf
	}
)

func NewUserCtrl(params UserCtrlParams) *UserCtrl {
	return &UserCtrl{
		userFacade: params.UserFacade,
	}
}

func (u *UserCtrl) SignUp(c *gin.Context) {
	funcName := "controller.(*userCtrl).SignUp"
	ctx := c.Request.Context()

	var req entity.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[%v] error bind request, err: %v\n", funcName, err)
		c.JSON(http.StatusBadRequest, entity.SignUpResponse{
			Message: err.Error(),
		})
		return
	}

	err := u.userFacade.SignUp(ctx, entity.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Username: req.Username,
	})
	if err != nil {
		log.Printf("[%v] error sign up, err: %v\n", funcName, err)
		c.JSON(http.StatusInternalServerError, entity.SignUpResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.SignUpResponse{
		Message: "success",
	})
}

func (u *UserCtrl) Login(c *gin.Context) {
	funcName := "controller.(*userCtrl).SignUp"
	ctx := c.Request.Context()

	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[%v] error bind request, err: %v\n", funcName, err)
		c.JSON(http.StatusBadRequest, entity.LoginResponse{
			Message: err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		log.Printf("[%v] error validate request, err: %v\n", funcName, err)
		c.JSON(http.StatusBadRequest, entity.LoginResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := u.userFacade.Login(ctx, req)
	if err != nil {
		log.Printf("[%v] error login, err: %v\n", funcName, err)
		c.JSON(http.StatusInternalServerError, entity.SignUpResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.LoginResponse{
		Message: "success",
		Token:   token,
	})
}
