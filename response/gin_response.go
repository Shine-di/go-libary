/**
 * @author: D-S
 * @date: 2020/8/10 6:50 下午
 */

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	RESULT = "result"
	ERROR  = "error"
)

func ResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		result, ok := c.Get(RESULT)
		if ok {
			c.JSON(http.StatusOK, result)
			return
		}
		err, ok := c.Get(ERROR)
		if ok {
			c.JSON(http.StatusOK, err)
			return
		}

	}
}

func Success(total int64, perPage int, result interface{}) ApiResponse {
	return ApiResponse{
		Code:    1,
		Message: "success",
		Total:   total,
		PerPage: int64(perPage),
		Result:  result,
	}
}

func Error(message string) ApiResponse {
	return ApiResponse{
		Code:    0,
		Message: message,
		Total:   0,
		PerPage: 0,
		Result:  nil,
	}
}

func responseOutput(c *gin.Context, code int, message string, total, perPage int64, result interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: message,
		Total:   total,
		PerPage: perPage,
		Result:  result,
	})
}

//{"match_id":817,"game_id":3,"series_id":137,"win_team":220,"data":"[{\"team_id\":225,\"team_score\":[8,8,0],\"kpr\":\"4.00\",\"kill\":70,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"\",\"death\":95,\"avg_assist\":\"\",\"assist\":19,\"avg_flash_assist\":\"0.00\",\"flash_assist\":4,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"0.00\",\"kd_diff\":\"-25.00\",\"kd_ratio\":\"0.00\",\"adr\":\"0.00\",\"kast\":\"0.00\",\"impact\":\"0.00\",\"rating\":\"0.00\",\"player\":[{\"player_id\":180,\"kpr\":\"0.79\",\"kill\":19,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.67\",\"death\":16,\"avg_assist\":\"0.00\",\"assist\":2,\"avg_flash_assist\":\"0.00\",\"flash_assist\":1,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"+2\",\"kd_diff\":\"+3\",\"kd_ratio\":\"1.1875\",\"adr\":\"64.1\",\"kast\":\"0.667\",\"impact\":\"1.10\",\"rating\":\"1.04\"},{\"player_id\":182,\"kpr\":\"0.62\",\"kill\":15,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.75\",\"death\":18,\"avg_assist\":\"0.00\",\"assist\":4,\"avg_flash_assist\":\"0.00\",\"flash_assist\":0,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"+3\",\"kd_diff\":\"-3\",\"kd_ratio\":\"0.8333333333333334\",\"adr\":\"85.4\",\"kast\":\"0.625\",\"impact\":\"1.20\",\"rating\":\"1.00\"},{\"player_id\":183,\"kpr\":\"0.67\",\"kill\":16,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.83\",\"death\":20,\"avg_assist\":\"0.00\",\"assist\":3,\"avg_flash_assist\":\"0.00\",\"flash_assist\":0,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"+3\",\"kd_diff\":\"-4\",\"kd_ratio\":\"0.8\",\"adr\":\"73.1\",\"kast\":\"0.708\",\"impact\":\"0.79\",\"rating\":\"0.88\"},{\"player_id\":179,\"kpr\":\"0.46\",\"kill\":11,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.83\",\"death\":20,\"avg_assist\":\"0.00\",\"assist\":6,\"avg_flash_assist\":\"0.00\",\"flash_assist\":2,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"-1\",\"kd_diff\":\"-9\",\"kd_ratio\":\"0.55\",\"adr\":\"72.9\",\"kast\":\"0.625\",\"impact\":\"0.75\",\"rating\":\"0.76\"},{\"player_id\":181,\"kpr\":\"0.38\",\"kill\":9,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.88\",\"death\":21,\"avg_assist\":\"0.00\",\"assist\":4,\"avg_flash_assist\":\"0.00\",\"flash_assist\":1,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"-3\",\"kd_diff\":\"-12\",\"kd_ratio\":\"0.42857142857142855\",\"adr\":\"55.4\",\"kast\":\"0.45799999999999996\",\"impact\":\"0.36\",\"rating\":\"0.46\"}]},{\"team_id\":220,\"team_score\":[16,7,9],\"kpr\":\"6.00\",\"kill\":95,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"\",\"death\":71,\"avg_assist\":\"\",\"assist\":15,\"avg_flash_assist\":\"0.00\",\"flash_assist\":2,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"0.00\",\"kd_diff\":\"24.00\",\"kd_ratio\":\"1.00\",\"adr\":\"0.00\",\"kast\":\"0.00\",\"impact\":\"0.00\",\"rating\":\"0.00\",\"player\":[{\"player_id\":330,\"kpr\":\"1.08\",\"kill\":26,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.58\",\"death\":14,\"avg_assist\":\"0.00\",\"assist\":3,\"avg_flash_assist\":\"0.00\",\"flash_assist\":0,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"+2\",\"kd_diff\":\"+12\",\"kd_ratio\":\"1.8571428571428572\",\"adr\":\"114.2\",\"kast\":\"0.875\",\"impact\":\"1.98\",\"rating\":\"1.80\"},{\"player_id\":332,\"kpr\":\"1.04\",\"kill\":25,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.58\",\"death\":14,\"avg_assist\":\"0.00\",\"assist\":5,\"avg_flash_assist\":\"0.00\",\"flash_assist\":0,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"-2\",\"kd_diff\":\"+11\",\"kd_ratio\":\"1.7857142857142858\",\"adr\":\"108.4\",\"kast\":\"0.792\",\"impact\":\"1.57\",\"rating\":\"1.57\"},{\"player_id\":329,\"kpr\":\"0.67\",\"kill\":16,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.62\",\"death\":15,\"avg_assist\":\"0.00\",\"assist\":2,\"avg_flash_assist\":\"0.00\",\"flash_assist\":1,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"0\",\"kd_diff\":\"+1\",\"kd_ratio\":\"1.0666666666666667\",\"adr\":\"64.2\",\"kast\":\"0.667\",\"impact\":\"1.26\",\"rating\":\"1.06\"},{\"player_id\":331,\"kpr\":\"0.54\",\"kill\":13,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.50\",\"death\":12,\"avg_assist\":\"0.00\",\"assist\":2,\"avg_flash_assist\":\"0.00\",\"flash_assist\":1,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"0\",\"kd_diff\":\"+1\",\"kd_ratio\":\"1.0833333333333333\",\"adr\":\"58.0\",\"kast\":\"0.708\",\"impact\":\"0.59\",\"rating\":\"0.97\"},{\"player_id\":328,\"kpr\":\"0.62\",\"kill\":15,\"avg_headshot\":\"0.00\",\"headshot\":0,\"dpr\":\"0.67\",\"death\":16,\"avg_assist\":\"0.00\",\"assist\":3,\"avg_flash_assist\":\"0.00\",\"flash_assist\":0,\"avg_entry_kill\":\"0.00\",\"entry_kill\":0,\"avg_entry_death\":\"0.00\",\"entry_death\":0,\"one_win_multi\":0,\"multi_kill\":0,\"fk_diff\":\"-4\",\"kd_diff\":\"-1\",\"kd_ratio\":\"0.9375\",\"adr\":\"62.1\",\"kast\":\"0.625\",\"impact\":\"0.80\",\"rating\":\"0.92\"}]}]"}
