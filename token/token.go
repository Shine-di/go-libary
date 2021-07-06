/**
 * @author: D-S
 * @date: 2020/9/8 3:37 下午
 */

package token

import (
	"github.com/dishine/libary/response"
	"github.com/gin-gonic/gin"
)

type GameAuth string


func TokenMiddleware(game GameAuth,tokenMap map[string][]GameAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			c.Set(response.ERROR, response.Error("token error"))
			c.Abort()
			return
		}
		b, ok := tokenMap[token]
		if !ok {
			c.Set(response.ERROR, response.Error("token error"))
			c.Abort()
			return
		}
		if !auth(game, b) {
			c.Set(response.ERROR, response.Error("auth error"))
			c.Abort()
			return
		}
		c.Next()
	}
}

func auth(game GameAuth, auths []GameAuth) bool {
	for _, e := range auths {
		if e == game {
			return true
		}
	}
	return false
}
