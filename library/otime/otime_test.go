package otime

import (
	"context"
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
