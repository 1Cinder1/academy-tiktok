package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	myjwt "go_tiktok/utils/jwt"
	"net/http"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  "not logged in",
				"ok":   false,
			})

			c.Abort()
			return
		}

		jwtConfig := global.Config.Middleware.Jwt
		j := myjwt.NewJWT(&myjwt.Config{
			SecretKey: jwtConfig.SecretKey,
		})

		mc, err := j.ParseToken(token)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  err.Error(),
				"ok":   false,
			})
			c.Abort()
			return
		}
		if mc.ExpiresAt.Unix()-time.Now().Unix() < mc.BufferTime && mc.ExpiresAt.Unix()-time.Now().Unix() > 0 {
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Middleware.Jwt.ExpiresTime) * time.Second))
			newToken, _ := j.GenerateToken(mc)
			newClaims, _ := j.ParseToken(newToken)
			err = global.Rdb.Set(c,
				fmt.Sprintf("jwt；%d", newClaims.BaseClaims.Id),
				newToken,
				time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
			if err != nil {
				global.Logger.Error("set redis key failed.",
					zap.Error(err),
					zap.String("key", "jwt:[id]"), zap.Int("id", newClaims.BaseClaims.Id),
				)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "internal err",
					"ok":   false,
				})
				c.Abort()
				return
			}
		}

		c.Set("id", mc.BaseClaims.Id)
		c.Next()
	}

}
