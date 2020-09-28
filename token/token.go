/**
 * @author: D-S
 * @date: 2020/9/8 3:37 下午
 */

package token

import (
	"github.com/Shine-di/go-libary/response"
	"github.com/gin-gonic/gin"
)

type gameAuth string

const (
	LOL      gameAuth = "lol"
	DOTA     gameAuth = "dota2"
	CSGO     gameAuth = "csgo"
	KOG      gameAuth = "kog"
	USER_ONE          = "zqcvN2BAjeao083yLZbVtOCSXDU4isJ57pKFh6PHmgIEdx9klw"

	TEST_USER_1 = "1iYjxTJHn7sN3eBqbmtzpklLaS65OVDvuR0UXrc8PFCMK"
	TEST_USER_2 = "YKHphjw0TEoqzsxgQeA1C5MavDc4lOmt6duy29b7S8RLG"
)

var (
	tokenMap = map[string][]gameAuth{
		USER_ONE: {
			LOL,
			DOTA,
			CSGO,
			KOG,
		},
		TEST_USER_1: {
			LOL,
		},
		TEST_USER_2: {
			LOL,
		},
	}
)

func TokenMiddleware(game gameAuth) gin.HandlerFunc {
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

func auth(game gameAuth, auths []gameAuth) bool {
	for _, e := range auths {
		if e == game {
			return true
		}
	}
	return false
}
