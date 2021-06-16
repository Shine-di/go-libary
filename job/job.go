/**
 * @author: D-S
 * @date: 2020/8/29 1:10 下午
 */

package job

import (
	"fmt"
	"github.com/dishine/libary/log"
	"github.com/robfig/cron/v3"
)

const (
	SecondPer1  = "*/1 * * * * ?"
	SecondPer5  = "*/5 * * * * ?"
	SecondPer10 = "*/10 * * * * ?"
	SecondPer15 = "*/15 * * * * ?"
	SecondPer30 = "*/30 * * * * ?"

	HourPer1 = "0 0 */1 * * ?"
	HourPer2 = "0 0 */2 * * ?"
	HourPer3 = "0 0 */3 * * ?"

	MinPer1  = "0 */1 * * * ?"
	MinPer2  = "0 */2 * * * ?"
	MinPer3  = "0 */3 * * * ?"
	MinPer4  = "0 */4 * * * ?"
	MinPer5  = "0 */5 * * * ?"
	MinPer6  = "0 */6 * * * ?"
	MinPer7  = "0 */7 * * * ?"
	MinPer10 = "0 */10 * * * ?"
	MinPer20 = "0 */20 * * * ?"
	MinPer30 = "0 */30 * * * ?"

	Timing2310 = "0 10 23 * * ?"
	Timing2330 = "0 30 23 * * ?"
	Timing2200 = "0 0 22 * * ?"
	Timing2300 = "0 0 23 * * ?"
	Timing0100 = "0 0 1 * * ?"
	Timing0200 = "0 0 2 * * ?"
	Timing0300 = "0 0 3 * * ?"
	Timing0400 = "0 0 4 * * ?"
	Timing0500 = "0 0 5 * * ?"
	Timing0600 = "0 0 6 * * ?"
	Timing0700 = "0 0 7 * * ?"
)

func InitJobs(jobs []SJob, b bool) {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)
	for _, e := range jobs {
		if b {
			go e.Run()
		}
		EntryID, err := c.AddJob(e.GetSpec(), e)
		if err != nil {
			log.Info("添加定时任务失败")
			continue
		}
		log.Info(fmt.Sprintf("添加定时任务成功: --%v--", EntryID))
	}
	c.Run()
	select {}
}

type SJob interface {
	GetSpec() string
	cron.Job
}

type Job interface {
	GetSpec() string
	DoFirst() bool
	cron.Job
}

type Func struct {
	t    string
	f    bool
	Func func()
}

func NewFunc(t string, f func(), doFirst bool) *Func {
	return &Func{
		t:    t,
		f:    doFirst,
		Func: f,
	}
}

func (f *Func) GetSpec() string {
	return f.t
}
func (f *Func) DoFirst() bool {
	return f.f
}

func (f *Func) Run() {
	f.Func()
}
