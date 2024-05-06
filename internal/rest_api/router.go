package rest_api

import (
	"coincap/internal/initmodule"
	"coincap/internal/rest_api/middleware"
	"coincap/pkg/cfg"

	"github.com/gin-gonic/gin"
)

func Router(dep *initmodule.RestAPIWrapper, config *cfg.ConfigSchema) *gin.Engine {
	ctrl := dep.Controller

	r := gin.Default()

	r.GET("/ping", ctrl.Ping)
	v1 := r.Group("/v1")
	v1.Use(
		middleware.IsAuthenticated(config, dep.Usecases.UserUsecaseItf),
	)
	{
		r.PUT("v1/user/login", ctrl.Login)
		r.PUT("v1/user/signup", ctrl.SignUp)
		user := v1.Group("/user")
		user.PUT("/logout")
	}
	{
		mycoin := v1.Group("/mycoin")
		mycoin.GET("/stared", ctrl.GetStaredCoin)
		mycoin.POST("/add", ctrl.AddStaredCoin)
		mycoin.DELETE("/delete", ctrl.DeleteStaredCoin)
	}
	{
		coin := v1.Group("/coin")
		coin.GET("/list", ctrl.GetCoins)
	}

	return r
}
