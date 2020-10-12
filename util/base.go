/**
 * @author: D-S
 * @date: 2020/8/17 7:17 下午
 */

package util

import (
	"strings"
	"time"
)

const TimeFormat = "2006-01-02T15:04:05Z"

// 1 未开始 2 进行中 ，3 已结束
func GetStatus(beginAt, endAt string) int64 {
	now := time.Now().Unix()
	if beginAt == "" {
		return 1
	}
	if beginAt != "" && endAt == "" {
		beginTime, err := time.ParseInLocation(TimeFormat, beginAt, time.UTC)
		if err != nil {
			return 1
		}
		if now > beginTime.Unix() {
			return 2
		} else {
			return 1
		}
	}
	if beginAt != "" && endAt != "" {
		beginTime, err := time.ParseInLocation(TimeFormat, beginAt, time.UTC)
		if err != nil {
			return 1
		}
		endTime, err := time.ParseInLocation(TimeFormat, endAt, time.UTC)
		if err != nil {
			return 1
		}
		if beginTime.Unix() > now {
			return 1
		} else {
			if now < endTime.Unix() {
				return 2
			} else {
				return 3
			}
		}
	}
	return 1
}

func Deduplication(data []int64) []int64 {
	Map := make(map[int64]bool, 0)
	for _, v := range data {
		Map[v] = true
	}
	result := make([]int64, 0)
	for k := range Map {
		result = append(result, k)
	}
	return result
}

//替换特殊字符串
func ReplaceSpecial(s string) string {
	if strings.Contains(s, "\\") {
		s = strings.ReplaceAll(s, "\\", "")
	}
	if strings.Contains(s, "\r") {
		s = strings.ReplaceAll(s, "\r", "")
	}
	if strings.Contains(s, "\n") {
		s = strings.ReplaceAll(s, "\n", "")
	}
	if strings.Contains(s, "\t") {
		s = strings.ReplaceAll(s, "\t", "")
	}
	return s
}
