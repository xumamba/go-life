package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

func func1() error {
	respChain := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		respChain <- 1
		close(respChain)
	}()

	select {
	case resp := <- respChain:
		fmt.Println("Receive response: ", resp)
		return nil
	case <-time.After(time.Second * 2):
		return errors.New("timeout")
	}

}

func func2(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	respChain := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		respChain <- 1
	}()
	select {
	case <- ctx.Done():
		fmt.Println("cancel")
		return errors.New("cancel")
	case resp :=  <- respChain:
		fmt.Println(resp)
		return nil
	}
}

func func3(ctx context.Context) {
	hctx, cancelFunc := context.WithTimeout(ctx, time.Second*5)
	defer cancelFunc()

	resp := make(chan struct{}, 1)
	go func() {
		time.Sleep(time.Second * 6)
		resp <- struct{}{}
	}()

	select {
	case <- hctx.Done():
		fmt.Println("timeout ctx")
		fmt.Println(hctx.Err())
	case v1 := <-resp:
		fmt.Println("function handle done")
		fmt.Printf("result: %v\n", v1)
	}
	fmt.Println("test2 finish")
	return
}

func TestFunc1(t *testing.T) {
	err := func1()
	fmt.Println("func1: ", err)

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		err := func2(ctx, wg)
		fmt.Println("func2: ", err)
	}()
	time.Sleep(time.Second * 4)
	cancelFunc()
	wg.Wait()

}