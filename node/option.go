/**
 * @author: D-S
 * @date: 2021/1/25 11:45 上午
 */

package node

import "github.com/dishine/libary/job"

type OptionFunc interface {
	GetOrder() Order
	GetOptionFunc() Func
}

type Options struct {
	Jobs    []job.Job // 定时任务
	Options []OptionFunc
}
