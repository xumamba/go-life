package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func worker(i int, ch chan bool, wg *sync.WaitGroup)  {
	fmt.Println(i, "goroutine count = ", runtime.NumGoroutine())
	<- ch
	wg.Done()
}


func TestGoroutineControl(t *testing.T) {
	ch := make(chan bool, 3)
	wg := sync.WaitGroup{}
	for i := 0; i<10; i++{
		wg.Add(1)
		ch <- true
		go worker(i, ch, &wg)
	}
	wg.Wait()
}

func worker2(ch chan int, wg *sync.WaitGroup)  {
	for task := range ch{
		fmt.Println("task: ", task, "goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func producer(task int, ch chan int, wg *sync.WaitGroup)  {
	wg.Add(1)
	ch <- task
}

func TestGoroutineControl2(t *testing.T)  {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	
	for i := 0;i <3;i ++{
		go worker2(ch, &wg)
	}
	
	for i := 0; i < 10; i++{
		producer(i, ch, &wg)
	}
	wg.Wait()
}

func TestChar(t *testing.T) {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}

func numJewelsInStones(jewels string, stones string) int {
	if len(jewels) == 0{
		return 0
	}
	result := map[rune]uint32{}
	r := 0
	for _, j := range jewels{
		result[j] = 0
	}
	for _, c := range stones{
		if _, ok := result[c]; ok{
			result[c]++
			r++
		}
	}
	fmt.Println(result)
	return r
}