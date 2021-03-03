package otime

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeStamp_Scan(t *testing.T) {
	var (
		testDuration Duration
		err          error
		testCtx      = context.Background()
	)
	err = testDuration.UnmarshalText([]byte("2s"))
	assert.Nil(t, err)
	duration, ctx, cancelFunc := testDuration.Shrink(testCtx)
	defer cancelFunc()
	assert.Equal(t, 2*time.Second, time.Duration(duration))
	deadline, ok := ctx.Deadline()
	assert.True(t, ok)
	assert.True(t, 2*time.Second > time.Until(deadline))
	assert.True(t, time.Second < time.Until(deadline))
}

func TestFunc(t *testing.T) {
	ts1, ts2 := int64(1613680159), int64(1614763332)
	days := GetDaysByTimestamp(ts1, ts2)
	fmt.Println(days)

	fmt.Println(GetDate(ts1))
	fmt.Println(GetDate(ts2))

}
