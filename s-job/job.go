/**
 * @author: D-S
 * @date: 2020/8/29 1:10 下午
 */

package s_job

import (
	"fmt"
	"github.com/Shine-di/go-libary/log"
	"github.com/robfig/cron/v3"
)

const (
	Second_5  = "*/5 * * * * ?"
	Second_10 = "*/10 * * * * ?"
	Second_30 = "*/30 * * * * ?"
	Hour_1    = "0 0 */1 * * ?"

	Min_1  = "0 */1 * * * ?"
	Min_3  = "0 */3 * * * ?"
	Min_5  = "0 */5 * * * ?"
	Min_10 = "0 */10 * * * ?"
	Min_20 = "0 */20 * * * ?"
	Min_30 = "0 */30 * * * ?"

	Hour_15    = "0 10 23 * * ?"
	Hour_23_30 = "0 30 23 * * ?"
	Hour_22    = "0 0 22 * * ?"
	Hour_23    = "0 0 23 * * ?"
	Hour_24    = "0 0 0 1 * ?"
	Hour_1_P   = "0 0 1 * * ?"
	Hour_2_P   = "0 0 2 * * ?"
	Hour_3_P   = "0 0 3 * * ?"
	Hour_4_P   = "0 0 4 * * ?"
	Hour_5_P   = "0 0 5 * * ?"
	Hour_6_P   = "0 0 6 * * ?"
	Hour_7_P   = "0 0 7 * * ?"
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
