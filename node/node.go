/**
 * @author: D-S
 * @date: 2021/1/25 11:43 上午
 */

package node

import (
	"errors"
	"fmt"
	"github.com/dishine/libary/job"
	"github.com/dishine/libary/log"
	"github.com/judwhite/go-svc/svc"
	"github.com/robfig/cron/v3"
)

type Order string

// 需要执行的 func
type OptionFunc func() error

const (
	Before Order = "before"
	After  Order = "after"
)

//节点启动的 流程
type Node struct {
	Jobs       []job.Job // 定时任务
	NodeName   string
	Options    []Option
	beforeFunc []OptionFunc
	afterFunc  []OptionFunc
}

// 默认参数
func (n *Node) defaultNode() Node {
	return Node{
		Jobs:       []job.Job{},
		NodeName:   "",
		Options:    nil,
		beforeFunc: []OptionFunc{},
		afterFunc:  []OptionFunc{},
	}
}

func (n *Node) Init(env svc.Environment) error {
	if n.Options != nil && len(n.Options) > 0 {
		for _, e := range n.Options {
			switch e.GetOrder() {
			case Before:
				n.beforeFunc = append(n.beforeFunc, e.GetOptionFunc())
			case After:
				n.afterFunc = append(n.afterFunc, e.GetOptionFunc())
			default:
				return errors.New("option func order err")
			}
		}
	}
	log.InitLogger(n.NodeName)
	return nil
}

func (n *Node) Start() error {
	// 执行之前任务
	if len(n.beforeFunc) > 0 {
		for _, f := range n.beforeFunc {
			if err := f(); err != nil {
				return err
			}
		}
	}
	// 执行之后的任务
	if len(n.afterFunc) > 0 {
		for _, f := range n.afterFunc {
			if err := f(); err != nil {
				return err
			}
		}
	}
	// 执行定时任务
	if len(n.Jobs) > 0 {
		c := cron.New(
			cron.WithSeconds(),
			cron.WithChain(cron.Recover(cron.DefaultLogger)),
		)
		for _, e := range n.Jobs {
			if e.DoFirst() {
				go e.Run()
			}
			EntryID, err := c.AddJob(e.GetSpec(), e)
			if err != nil {
				return err
			}
			log.Info(fmt.Sprintf("添加定时任务成功: --%v--", EntryID))
		}
		go c.Run()
	}
	return nil
}

func (n *Node) Stop() error {
	log.Info("停止服务.......")
	return nil
}
