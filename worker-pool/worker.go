package worker_pool

import (
	"errors"
	"fmt"
)

var (
	ErrBadTask    = errors.New("bad task")
	ErrBadSize    = errors.New("bad size")
	ErrNotStarted = errors.New("not started")
)

type Task func() error

type Worker struct {
	ID         int
	taskQueue  chan Task
	resultChan chan<- error
}

func NewWorker(ID int, taskQueue chan Task, resultChan chan<- error) *Worker {
	return &Worker{
		ID:         ID,
		taskQueue:  taskQueue,
		resultChan: resultChan,
	}
}

func (wr *Worker) Start() {
	fmt.Printf("Starting worker %d\n", wr.ID)

	go func() {
		//fmt.Printf("wp.Start(ID: %d) taskQue len: %d\n", wr.ID, len(wr.taskQueue))
		for task := range wr.taskQueue {
			//fmt.Printf("wp.Start(ID: %v) received task: %d\n", wr.ID, task)
			err := task()
			//fmt.Printf("task(ID: %v) result: %v\n", task, err)
			if err != nil {
				//fmt.Printf("sending result to resultChan: %v\n", err)
				//fmt.Printf("wp.Start(ID: %d) resultChan len: %d\n", wr.ID, len(wr.resultChan))
				wr.resultChan <- err
			}
		}
	}()
}
