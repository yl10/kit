# 时间小工具

- DateRange 获取指定间隔的日期 slice
```go
// 以 from 为起始时间, 以 d 为间隔, 获取 num 个时间
func DateRange(from time.Time, d time.Duration, num int) []time.Time
```

- DayRange 获取当天的起止时间 00:00:00 - 23:59:59
```go
func DayRange(dt time.Time) (start time.Time, end time.Time)
```
