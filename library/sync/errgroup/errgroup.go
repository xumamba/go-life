package errgroup

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

type WaitGroup struct {
	err        error
	wg         sync.WaitGroup
	errOnce    sync.Once
	workerOnce sync.Once
	ch         chan func() error
	fs         []func() error
	cancel     func()
}

func WithContext(ctx context.Context) (*WaitGroup, context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	return &WaitGroup{cancel: cancel}, childCtx
}

func (wg *WaitGroup) do(f func() error) {
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
	err = f()
}

func (wg *WaitGroup) GOMAXPROCS(m int) {
	if m <= 0 {
		panic("m can not later than 0")
	}
	wg.workerOnce.Do(func() {
		wg.ch = make(chan func() error, m)
		for i := 0; i < m; i++ {
			go func() {
				for f := range wg.ch {
					wg.do(f)
				}
			}()
		}
	})
}

func (wg *WaitGroup) Go(f func() error) {
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
