package worker_pool

//
//import (
//	"errors"
//	"fmt"
//	"sync"
//)
//
//var (
//	ErrBadTask    = errors.New("bad task")
//	ErrBadSize    = errors.New("bad size")
//	ErrNotStarted = errors.New("not started")
//)
//
//type Task func() error
//
//type WorkerPool struct {
//	started bool
//
//	numWorkers int
//	tasks      chan Task
//	results    chan error
//}
//
//func NewWorkerPool(size int) (*WorkerPool, error) {
//	if size <= 0 {
//		return nil, ErrBadSize
//	}
//
//	tasks := make(chan Task)
//	results := make(chan error)
//
//	fmt.Printf("tasks: %v\n", tasks)
//	fmt.Printf("results: %v\n", results)
//	return &WorkerPool{
//		started:    false,
//		numWorkers: size,
//		tasks:      tasks,
//		results:    results,
//	}, nil
//}
//
//func (wp *WorkerPool) Run() {
//
//	fmt.Println("started:", wp.started)
//
//	wp.started = true
//	fmt.Println("started:", wp.started)
//
//	var wg sync.WaitGroup
//
//	for w := 1; w <= wp.numWorkers; w++ {
//		wg.Add(1)
//
//		go func() {
//			defer wg.Done()
//			wp.worker(w)
//		}()
//	}
//
//	wg.Wait()
//
//}
//
//func (wp *WorkerPool) AddTask(task Task) error {
//	fmt.Println("AddTask() entered:", task)
//	if task == nil {
//		return ErrBadTask
//	}
//	if !wp.started {
//		return ErrNotStarted
//	}
//
//	select {
//	case wp.tasks <- task:
//		fmt.Printf("adding task %v\n", task)
//		return nil
//	default:
//		return nil
//	}
//
//}
//
//func (wp *WorkerPool) Result() chan error {
//	if !wp.started {
//		fmt.Println("Result(): ", wp.started)
//		return nil
//	}
//
//	fmt.Println("wp.results: ", wp.results)
//
//	return wp.results
//}
//
//func (wp *WorkerPool) worker(id int) {
//
//	fmt.Println("worker", id, "started")
//	fmt.Println("tasks len", len(wp.tasks))
//	for task := range wp.tasks {
//		fmt.Println("worker", id, "started  job", task)
//		err := task()
//		if err != nil {
//			wp.results <- err
//		}
//		fmt.Println("worker", id, "finished job", task)
//	}
//	fmt.Println("leawing worker")
//}
