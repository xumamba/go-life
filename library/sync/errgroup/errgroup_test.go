package errgroup

import (
	"context"
	"errors"
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithCancel(t *testing.T) {
	waitGroup := WithCancel(nil)
	assert.NotNil(t, waitGroup.ctx)
}

func TestWaitGroup(t *testing.T) {
	group := WaitGroup{}
	begin := time.Now()
	group.Go(sleep)
	group.Go(sleep)
	group.Go(sleep)
	group.Go(sleep)
	group.Wait()
	spend := math.Round(time.Since(begin).Seconds())
	t.Log("spend time:", spend)
	if spend != 1 {
		t.FailNow()
	}

	group2 := WithContext(context.Background())
	group2.GOMAXPROCS(2)
	begin2 := time.Now()
	group2.Go(sleep)
	group2.Go(sleep)
	group2.Go(sleep)
	group2.Go(sleep)
	group2.Wait()
	spend2 := math.Round(time.Since(begin2).Seconds())
	t.Log("spend time:", spend2)
	if spend2 != 2 {
		t.FailNow()
	}

	var doneErr error
	group3 := WithCancel(context.Background())
	group3.GOMAXPROCS(2)
	begin3 := time.Now()
	group3.Go(func(ctx context.Context) error {
		// time.Sleep(time.Second*2)
		return errors.New("mock an error")
	})
	group3.Go(func(ctx context.Context) error {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			doneErr = context.Canceled
		default:
			t.Log("should not be executed")
		}
		return nil
	})
	group3.Wait()
	spend3 := math.Round(time.Since(begin3).Seconds())
	t.Log("spend time:", spend3)
	if spend3 != 1 {
		t.FailNow()
	}
	if doneErr != context.Canceled {
		t.FailNow()
	}
}

func sleep(context.Context) error {
	time.Sleep(time.Second)
	return nil
}

var queue chan int

func TestWithContext(t *testing.T) {
	InitQueue()
	id := 0
	for {
		// 接收请求
		if queue != nil {
			select {
			case queue <- id:
			default:
				// 熔断、限流、降级处理
			}
		}
		time.Sleep(2 * time.Second)
		id++
	}
}

func DoSomething(id int) {
	var err error
	defer func() {
		if pErr := recover(); pErr != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("occur an panic：%s\n,stack:%s", pErr, buf)
		}
	}()
	fmt.Println(id, err)
}

func InitQueue() {
	queue = make(chan int, 64)
	go func() {
		for f := range queue {
			DoSomething(f)
		}
	}()
}
