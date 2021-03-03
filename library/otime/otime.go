package otime

import (
	"context"
	"database/sql/driver"
	"strconv"
	"time"
)

type TimeStamp int64

func (ts *TimeStamp) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case time.Time:
		*ts = TimeStamp(sc.Unix())
	case string:
		var parseInt int64
		parseInt, err = strconv.ParseInt(sc, 10, 64)
		*ts = TimeStamp(parseInt)
	}
	return
}

func (ts *TimeStamp) Value() (driver.Value, error) {
	return ts.Time(), nil
}

func (ts *TimeStamp) Time() time.Time {
	return time.Unix(int64(*ts), 0)
}

type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	duration, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(duration)
	}
	return err
}

func (d Duration) Shrink(ctx context.Context) (Duration, context.Context, context.CancelFunc) {
	if deadline, ok := ctx.Deadline(); ok {
		if ctxTimeout := time.Until(deadline); ctxTimeout < time.Duration(d) {
			return Duration(ctxTimeout), ctx, func() {}
		}
	}
	context, cancelFunc := context.WithTimeout(ctx, time.Duration(d))
	return d, context, cancelFunc
}

// 获取两个时间戳相隔天数
func GetDaysByTimestamp(ts1, ts2 int64) int {
	nowDate, _ := time.Parse("2006-01-02", time.Unix(ts2, 0).Format("2006-01-02"))
	aimDate, _ := time.Parse("2006-01-02", time.Unix(ts1, 0).Format("2006-01-02"))
	hours := nowDate.Sub(aimDate)
	days := int(hours.Hours() / 24)
	if days < 0 {
		days = -days
	}
	return days
}

// 获取指定时间戳对应的日期
func GetDate(timestamp int64) int {
	date, _ := strconv.Atoi(time.Unix(timestamp, 0).Format("20060102"))
	return date
}
