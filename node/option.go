/**
 * @author: D-S
 * @date: 2021/1/25 11:45 上午
 */

package node

import "github.com/dishine/libary/job"

type Options struct {
	Options  []OptionFunc
	HttpPort string
	Router   *Router
	Jobs     []job.Job // 定时任务
}

type OptionFunc interface {
	GetOrder() Order
	GetOptionFunc() Func
}

type DefaultOptionFunc struct {
	order      Order
	optionFunc Func
}

func NewOptionFunc(order Order, optionFunc Func) *DefaultOptionFunc {
	return &DefaultOptionFunc{
		order:      order,
		optionFunc: optionFunc,
	}
}

func (o *DefaultOptionFunc) GetOrder() Order {
	return o.order
}
func (o *DefaultOptionFunc) GetOptionFunc() Func {
	return o.optionFunc
}
