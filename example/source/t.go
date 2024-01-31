package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Task struct {
	Steps []*Step
}

func NewTask(steps ...*Step) *Task {
	return &Task{
		Steps: steps,
	}
}

func (t *Task) Run() {
	var wg = &sync.WaitGroup{}
	for _, v := range t.Steps {
		v.run(wg)
	}
	wg.Wait()
}

type Step struct {
	Name string
	deps []*Step
	done chan struct{}
	fn   func()
}

func NewStep(name string, fn func(), tasks ...*Step) *Step {
	task := &Step{
		Name: name,
		deps: tasks,
		done: make(chan struct{}),
		fn:   fn,
	}

	return task
}

func (task *Step) run(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(task.done)

		// Create a wait group to wait for all dependencies to close
		if len(task.deps) == 0 {
			task.fn()
			return
		}

		var depWG sync.WaitGroup

		// Increment the dependency wait group for each dependency
		for _, dep := range task.deps {
			depWG.Add(1)
			go func(dep *Step) {
				defer func() {
					if r := recover(); r != nil {
						log.Println(r)
					}
					depWG.Done()
				}()

				<-dep.done
			}(dep)
		}

		// Wait for all dependencies to close
		depWG.Wait()

		// All dependencies are closed, proceed with the task
		task.fn()
	}()

}

func main() {
	task1 := NewStep("task1", func() {
		fmt.Println("task1")
	})

	task4 := NewStep("task4", func() {
		time.Sleep(time.Second)
		fmt.Println("task4")
	}, task1)

	task2 := NewStep("task2", func() {
		fmt.Println("task2")
	}, task1)

	task3 := NewStep("task3", func() {
		fmt.Println("task3")
	}, task4, task2)

	task := NewTask(task1, task2, task3, task4)
	task.Run()
}
