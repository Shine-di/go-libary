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
	LOL  gameAuth = "lol"
	DOTA gameAuth = "dota2"
	CSGO gameAuth = "csgo"
	KOG  gameAuth = "kog"

	RATE gameAuth = "rate"
	// 自己测试用的token
	TEST = "VRI5y04dNlUJDYaoKzgHfmEwbBtuhvX7FrMOk3QjxeZAW8cSspLTCnq"
	// 商户 1
	USER_ONE = "D67lL1AYF2j9bXMy80RNaTU43fIQpPOzxrmscKdG5SvoCBtEHhZnwgJ"
	//USER_ONE = "zqcvN2BAjeao083yLZbVtOCSXDU4isJ57pKFh6PHmgIEdx9klw"
	// 商户测试1
	TEST_USER_1 = "TBoKrJFIkR87uxg6zawUADX1lH5tci9QS3fMjneWmEdZs2GNqOpL4P0"
	//TEST_USER_1 = "1iYjxTJHn7sN3eBqbmtzpklLaS65OVDvuR0UXrc8PFCMK"
	// 商户测试2
	TEST_USER_2 = "krulfOFQIHEiMhGCYW21tTZ3XgjNBsUS4xezonwVAJPK8pLRdq7Dav0"
	//TEST_USER_2 = "YKHphjw0TEoqzsxgQeA1C5MavDc4lOmt6duy29b7S8RLG"
)

var (
	tokenMap = map[string][]gameAuth{
		TEST: {
			LOL,
			DOTA,
			CSGO,
			KOG,
			RATE,
		},
		USER_ONE: {
			LOL,
			DOTA,
			CSGO,
			KOG,
		},
		TEST_USER_1: {
			LOL,
			DOTA,
			CSGO,
			KOG,
		},
		TEST_USER_2: {
			LOL,
			DOTA,
			CSGO,
			KOG,
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
