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
type Func func() error

const (
	Before Order = "before"
	After  Order = "after"
)

//节点启动的 流程
type Node struct {
	NodeName   string
	options    *Options
	beforeFunc []Func
	afterFunc  []Func
}

func New(name string, option *Options) *Node {
	if option == nil {
		return defaultNode(name)
	}
	return &Node{
		NodeName: name,
		options:  option,
	}
}

// 默认参数
func defaultNode(name string) *Node {
	return &Node{
		options: &Options{
			Options:  nil,
			HttpPort: "",
			Router:   nil,
			Jobs:     []job.Job{},
		},
		NodeName: name,
	}
}

func (n *Node) Init(env svc.Environment) error {
	if n.options.HttpPort != "" && n.options.Router == nil {
		return errors.New("web api param err: router is nil")
	}
	if n.options.Options != nil && len(n.options.Options) > 0 {
		for _, e := range n.options.Options {
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
	// 启动http
	if n.options.HttpPort != "" {
		go func() {
			if err := n.options.Router.Engine.Run(); err != nil {
				panic(fmt.Sprintf("run web api err: %v", err.Error()))
			}
		}()
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
	if len(n.options.Jobs) > 0 {
		c := cron.New(
			cron.WithSeconds(),
			cron.WithChain(cron.Recover(cron.DefaultLogger)),
		)
		for _, e := range n.options.Jobs {
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
