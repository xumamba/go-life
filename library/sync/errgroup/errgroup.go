package errgroup

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type WaitGroup struct {
	err error

	wg         sync.WaitGroup
	errOnce    sync.Once
	workerOnce sync.Once

	ch chan func(ctx context.Context) error
	fs []func(ctx context.Context) error

	ctx    context.Context
	cancel func()
}

func WithContext(ctx context.Context) *WaitGroup {
	if ctx == nil {
		ctx = context.Background()
	}
	return &WaitGroup{ctx: ctx}
}

func WithCancel(parentCtx context.Context) *WaitGroup {
	if parentCtx == nil {
		parentCtx = context.Background()
	}
	ctx, cancelFunc := context.WithCancel(parentCtx)
	return &WaitGroup{ctx: ctx, cancel: cancelFunc}
}

func (wg *WaitGroup) do(f func(ctx context.Context) error) {
	var err error
	defer func() {
		if pErr := recover(); pErr != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("occur an panic：%s\n,stack:%s", pErr, buf)
		}
		if err != nil {
			wg.errOnce.Do(func() {
				wg.err = err
				if wg.cancel != nil {
					wg.cancel()
				}
			})
		}
		wg.wg.Done()
	}()
	err = f(wg.ctx)
}

func (wg *WaitGroup) GOMAXPROCS(m int) {
	if m <= 0 {
		panic("m can't be less than 0")
	}
	wg.workerOnce.Do(func() {
		wg.ch = make(chan func(ctx context.Context) error, m)
		for i := 0; i < m; i++ {
			go func() {
				for f := range wg.ch {
					wg.do(f)
				}
			}()
		}
	})
}

func (wg *WaitGroup) Go(f func(ctx context.Context) error) {
	wg.wg.Add(1)
	if wg.ch != nil {
		select {
		case wg.ch <- f:
		default:
			wg.fs = append(wg.fs, f)
		}
		return
	}
	go wg.do(f)
}

func (wg *WaitGroup) Wait() error {
	if wg.ch != nil {
		for _, f := range wg.fs {
			wg.ch <- f
		}
	}
	wg.wg.Wait()
	if wg.ch != nil {
		close(wg.ch)
	}
	if wg.cancel != nil {
		wg.cancel()
	}
	return wg.err
}
