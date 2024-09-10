package main

import (
	worker_pool "design-patterns/workerpool/worker-pool"
	"errors"
	"fmt"
	"time"
)

func main() {
	shortTaskOk := func() error {
		fmt.Println("shortTaskOk")
		time.Sleep(1 * time.Second)
		return nil
	}

	longTaskOk := func() error {
		fmt.Println("longTaskOk")
		time.Sleep(15 * time.Second)
		return nil
	}

	shortTaskError := func() error {
		fmt.Println("shortTaskError")
		time.Sleep(1 * time.Second)
		return errors.New(time.Now().Format("2006-01-02 15:04:05"))
	}
	longTaskError := func() error {
		fmt.Println("longTaskError")
		time.Sleep(15 * time.Second)
		return errors.New(time.Now().Format("2006-01-02 15:04:05"))
	}

	wp, _ := worker_pool.NewWorkerPool(3)
	wp.Run()

	go func() {
		for elem := range wp.Results() {
			fmt.Printf("RESULT received: %v\n", elem)
		}
	}()

	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskOk)
	wp.AddTask(nil)
	wp.AddTask(longTaskOk)
	wp.AddTask(shortTaskError)
	wp.AddTask(shortTaskError)
	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskOk)
	wp.AddTask(longTaskError)
	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskError)
	wp.AddTask(shortTaskOk)
	wp.AddTask(shortTaskError)
	wp.AddTask(shortTaskOk)

	time.Sleep(20 * time.Second)
}
