/**
 * @author: D-S
 * @date: 2020/9/8 3:37 下午
 */

package token

import (
	"github.com/dishine/libary/response"
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
	TEST = "G5WaoTMYDuINAMVIQVmjANVhzCgDVT4AxPL9fAmMmSRCP6N8gE"

	// 正式的商户1 海南
	N_TEST_USER_6 = "2e5Xp9V2vuSMY7sfNR9QiwmIO8xxjCNqnJJzSSCzJd1kCaKSzn"

	N_TEST_USER_1 = "f1sv9BLGDUPXyeIdYuY0mgRfcthXifFGLVB14vVwniVc79NndV"
	N_TEST_USER_2 = "v5vB3h74Tptd9k16Gs84qaeQQMuQYqOk7SrA3Dfu5Lsm7v7bb9"
	N_TEST_USER_3 = "orxSUKhb7pLg4drzLMXV9TMmRsOCk1KO2vnz4ZppmBlnwfHqmf"
	N_TEST_USER_4 = "q8vQZXx3xeds35Zj3uaewQFjpsw8BUWcOHKqlhVBaG1N9mnaPC"
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
		//USER_ONE: {
		//	//LOL,
		//	//DOTA,
		//	//CSGO,
		//	//KOG,
		//},
		//TEST_USER_1: {
		//	//LOL,
		//	//DOTA,
		//	//CSGO,
		//	//KOG,
		//},
		//TEST_USER_2: {
		//	//LOL,
		//	//DOTA,
		//	//CSGO,
		//	//KOG,
		//},
		//
		N_TEST_USER_1: {
			//RATE,
		},
		N_TEST_USER_2: {
			//RATE,
		},
		N_TEST_USER_3: {
			//RATE,
		},
		N_TEST_USER_4: {
			//RATE,
		},
		N_TEST_USER_6: {
			RATE,
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
