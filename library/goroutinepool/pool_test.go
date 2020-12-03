package goroutinepool

import (
	"testing"
)

func TestGPool(t *testing.T)  {
	manager := NewManager(10)
	manager.StartWorkerPool()
}
