package dateutil

import (
	"time"
)

// DateRange 获取时间区间
// 以 from 为起始时间, 以 d 为间隔, 获取 num 个时间
func DateRange(from time.Time, d time.Duration, num int) []time.Time {
	r := make([]time.Time, 0)
	var t = from
	for i := 0; i < num; i++ {
		t = t.Add(d)
		r = append(r, t)
	}
	return r
}

// DayRange 获取当天的起止时间 00:00:00 - 23:59:59
func DayRange(dt time.Time) (time.Time, time.Time) {
	start, _ := time.Parse("2006-01-02 15:04:05", dt.Format("2006-01-02 00:00:00"))
	end := start.Add(24 * time.Hour).Add(-1 * time.Second)
	return start, end
}
