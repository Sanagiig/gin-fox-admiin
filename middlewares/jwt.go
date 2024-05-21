package middlewares

import (
	"errors"
	"gin-one/message"
	"gin-one/model/common/response"
	"gin-one/service/system"
	"gin-one/utils/ctx"
	"gin-one/utils/jwt"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ctx.GetToken(c)
		if token == "" {
			response.FailWithMessage(message.AUTH_FAILED, "", c)
			c.Abort()
			return
		}

		if system.JWTServiceApp.IsBlock(token) {
			response.FailWithMessage(message.TOKEN_IS_BLOCK, "", c)
			c.Abort()
			return
		}

		j := jwt.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, jwt.TokenExpired) {
				response.FailWithMessage(message.TOKEN_IS_EXPIRED, err.Error(), c)
				ctx.ClearToken(c)
				c.Abort()
				return
			}
		}

		c.Set("claims", claims)
		c.Next()
	}
}
