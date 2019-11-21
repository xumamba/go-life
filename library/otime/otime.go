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
