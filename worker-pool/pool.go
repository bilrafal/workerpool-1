package worker_pool

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var wpInstance *WorkerPool

type WorkerPool struct {
	isRunning   bool
	taskQueue   chan Task
	resultChan  chan error
	workerCount int
}

func NewWorkerPool(workerCount int) (*WorkerPool, error) {
	fmt.Println("EnteredNewWorkerPool()")
	if workerCount < 1 {
		return nil, ErrBadSize
	}

	once.Do(func() {
		fmt.Println("NewWorkerPool() => once.Do()")
		wpInstance = &WorkerPool{
			taskQueue:   make(chan Task),
			resultChan:  make(chan error),
			workerCount: workerCount,
			isRunning:   false,
		}
	})
	return wpInstance, nil
}

func (wp *WorkerPool) Run() {
	wp.isRunning = true
	fmt.Println("entered wp.Run()")
	for i := 0; i < wp.workerCount; i++ {
		//fmt.Println("worker", i, "start")
		worker := NewWorker(i, wp.taskQueue, wp.resultChan)
		worker.Start()
	}

}

func (wp *WorkerPool) AddTask(task Task) error {
	//fmt.Println("Entered AddTask()")
	if task == nil {
		//fmt.Println("task == nil ")
		return ErrBadTask
	}
	if !wp.isRunning {
		return ErrNotStarted
	}
	wp.taskQueue <- task
	//fmt.Printf("wp.AddTask() task len: %d\n", len(wp.taskQueue))
	return nil
}

func (wp *WorkerPool) Results() <-chan error {
	//fmt.Println("Entered RESULTS()")

	tickChan := time.NewTicker(time.Millisecond * 500).C

	for {
		select {
		case <-tickChan:
			if wp.isRunning {
				return wp.resultChan
			}
		default:
		}
	}

}
