package middleware

import (
	"coincap/internal/constant"
	"coincap/internal/usecase"
	"coincap/pkg/cfg"
	"coincap/pkg/converter"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthenticated(config *cfg.ConfigSchema, userUC usecase.UserUsecaseItf) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := struct {
			Message string `json:"message"`
		}{}
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			res.Message = "missing authorization"
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		var claims jwt.MapClaims
		_, err := jwt.ParseWithClaims(authorization, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JWT.SecretKey), nil
		})
		if err != nil {
			res.Message = err.Error()
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		if val, ok := claims[constant.KeyTokenExpiryDate]; ok {
			if converter.ToInt64(val.(string)) < time.Now().Unix() {
				res.Message = "expired jwt token"
				c.JSON(http.StatusForbidden, res)
				c.Abort()
				return
			}
		}

		userID := converter.ToInt64(claims[constant.KeyTokenUserID])
		userData, err := userUC.GetUserByID(context.Background(), userID)
		if err != nil {
			res.Message = "invalid user data"
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		if !userData.IsLogedIn {
			res.Message = "please login first"
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		seed := claims[constant.KeyTokenRandomSeed].(string)
		email := claims[constant.KeyTokenEmail].(string)

		if seed != userData.Seed || email != userData.Email {
			res.Message = "stolen credential!"
			c.JSON(http.StatusForbidden, res)
			c.Abort()
			return
		}

		c.Next()
	}
}
