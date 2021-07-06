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
	//N_TEST_USER_6 = "2e5Xp9V2vuSMY7sfNR9QiwmIO8xxjCNqnJJzSSCzJd1kCaKSzn"

	N_TEST_USER_1 = "f1sv9BLGDUPXyeIdYuY0mgRfcthXifFGLVB14vVwniVc79NndV"
	N_TEST_USER_2 = "v5vB3h74Tptd9k16Gs84qaeQQMuQYqOk7SrA3Dfu5Lsm7v7bb9"
	N_TEST_USER_3 = "orxSUKhb7pLg4drzLMXV9TMmRsOCk1KO2vnz4ZppmBlnwfHqmf"
	N_TEST_USER_4 = "q8vQZXx3xeds35Zj3uaewQFjpsw8BUWcOHKqlhVBaG1N9mnaPC"

	//20210706 更新
	N_TEST_USER_5 = "t9uNipkLF3XYEyPh37s8MPUfTPQt9sE0jQKANmpiafD9AJBntT"
	N_TEST_USER_6 = "QwHdDkUY74VSauoHLNmsud4LVZJ8pzk8t1CiweoIV1illqZRfU"
	N_TEST_USER_7 = "4PXo7Bm4HQlLrOzI4RtuQbu616jjpqxJPWtME3lRqg3aqUdFJ9"
	N_TEST_USER_8 = "MqlqZ28IJTBw4o7llkOvdU33ulEiAdd7CscUDHUtH4h0FtrxI2"
	N_TEST_USER_9 = "3QRlMRyFK1yEAZhsnFrLZayVRjAHcSo1cvcKiWoJfixyxcrGEt"
)

var (
	tokenMap = map[string][]gameAuth{}
)
func init()  {
	tokenMap = map[string][]gameAuth{
		TEST: {
			LOL,
			DOTA,
			CSGO,
			KOG,
			RATE,
		},
		N_TEST_USER_1: {
			RATE,
		},
		//N_TEST_USER_2: {
		//	RATE,
		//},
		N_TEST_USER_3: {
			RATE,
		},
		N_TEST_USER_5: {
			RATE,
		},
		N_TEST_USER_6: {
			RATE,
		},
		N_TEST_USER_7: {
			RATE,
		},
		N_TEST_USER_8: {
			RATE,
		},
		N_TEST_USER_9: {
			RATE,
		},

	}
}

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

//func SyncToken( enum redis.PrefixEnum) error {
//	go func() {
//		for {
//			r := redis.GetRedis()
//			if r == nil {
//				log.Warn("连接redis错误")
//				continue
//			}
//			key := fmt.Sprintf("HTTP-APT-TOKEN")
//			d,err := r.GetValue(enum,key)
//			if err != nil {
//				log.Info("同步token错误",zap.Error(err))
//				continue
//			}
//
//			<- time.After(time.Minute)
//		}
//	}()
//	return nil
//}