package goroutinepool

import (
	"fmt"
)

type Manager struct {
	workerNum int
	workChan  chan *Worker
}

func (m *Manager) StartWorkerPool() {
	for i := 0; i < m.workerNum; i++ {
		w := &Worker{id: i}
		go w.Do(m.workChan)
	}
	m.KeepLiveWorkers()
}

func (m *Manager) KeepLiveWorkers() {
	for w := range m.workChan {
		fmt.Printf("Worker %d stopped with err: [%v] \n", w.id, w.err)
		w.err = nil
		go w.Do(m.workChan)
	}
}

func NewManager(num int) *Manager {
	return &Manager{
		workerNum: num,
		workChan:  make(chan *Worker, num),
	}
}