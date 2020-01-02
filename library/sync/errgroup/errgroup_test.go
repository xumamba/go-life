package errgroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCancel(t *testing.T) {
	waitGroup := WithCancel(nil)
	assert.NotNil(t, waitGroup.ctx)
}
