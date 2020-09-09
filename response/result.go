/**
 * @author: D-S
 * @date: 2020/9/9 10:44 上午
 */

package response

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Total   int64       `json:"total"`
	PerPage int64       `json:"per_page" form:"per_page"`
	Result  interface{} `json:"data"`
}
