package util

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// 获取当前时间的前一周开始时间和结束时间
func GetLastWeekTime(now time.Time) (int64, int64) {
	if now.Unix() == 0 {
		now = time.Now()
	}
	offset := int(now.Weekday()) // 今天是周几
	if offset == 0 {
		offset = 6
	} else {
		offset = offset - 1
	}

	weekEndTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, 6-offset)
	weekStartTime := weekEndTime.AddDate(0, 0, -7)
	return weekStartTime.Unix() + 1, weekEndTime.Unix()
}

func GetLastMonthTime(now time.Time) (int64, int64) {
	if now.Unix() == 0 {
		now = time.Now()
	}
	year := now.Year()
	month := now.Month() + 1

	if month == 13 {
		month = 1
		year = year + 1
	}

	monthEndTime := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Unix() - 1
	monthStartTime := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local).Unix()

	return monthStartTime, monthEndTime
}

func DecrByLua(ctx context.Context, c redis.Conn, key string, qty uint32) (bool, error) {
	if qty == 0 {
		qty = 1
	}

	script := `
	local current_value = tonumber(redis.call('GET', KEYS[1]))
	local qty = tonumber(ARGV[1])
	if current_value >= qty then
		redis.call('DECRBY', KEYS[1], ARGV[1])
		return true
	else
		return false
	end
	`
	return redis.NewScript(script).Run(ctx, c, []string{key}, qty).Bool()
}
