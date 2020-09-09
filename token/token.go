/**
 * @author: D-S
 * @date: 2020/9/8 3:37 下午
 */

package token

import (
	gin_response "github.com/Shine-di/go-libary/response"
	"github.com/gin-gonic/gin"
)

type gameAuth string

const (
	LOL      gameAuth = "lol"
	DOTA     gameAuth = "dota2"
	USER_ONE          = "zqcvN2BAjeao083yLZbVtOCSXDU4isJ57pKFh6PHmgIEdx9klw"
)

var (
	tokenMap = map[string][]gameAuth{
		USER_ONE: {
			LOL,
			DOTA,
		},
	}
)

func TokenMiddleware(game gameAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			gin_response.Error(c, "token error")
			c.Abort()
			return
		}
		b, ok := tokenMap[token]
		if !ok {
			gin_response.Error(c, "token error")
			c.Abort()
			return
		}
		if !auth(game, b) {
			gin_response.Error(c, "auth error")
			c.Abort()
			return
		}
		c.Next()
	}
}

func auth(game gameAuth, auths []gameAuth) bool {
	for _, e := range auths {
		if e == game {
			return true
		}
	}
	return false
}
