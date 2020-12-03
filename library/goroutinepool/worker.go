package goroutinepool

import (
	"fmt"
	"time"
)

type Worker struct {
	id  int
	err error
}

func (w *Worker) Do(workerChan chan<- *Worker) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				w.err = err
			} else {
				w.err = fmt.Errorf("panic happened with [%v]", r)
			}
		}
		workerChan <- w
	}()

	fmt.Println("Start Worker...ID = ", w.id)

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
	}

	panic("worker panic..")

	return err
}
