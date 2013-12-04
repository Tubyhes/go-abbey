package main

import (
	"fmt"
	"time"
)

type Component interface {
	init()
	process(x, y int)
	end() int
}

type Task struct {
	C *Component
	X, Y int
	TaskNum int
}

type Sum struct {
	result int
}

type Multiply struct {
	result int
}

func Process (x, y int, c *Component) int {
	(*c).init()
	(*c).process(x, y)
	return (*c).end()
}

func (s *Sum) init () {

}

func (s *Sum) process (x, y int) {
	s.result = x + y
}

func (s *Sum) end () int {
	return s.result
}

func (m *Multiply) init () {

}

func (m *Multiply) process (x, y int) {
	m.result = x * y
}

func (m *Multiply) end () int {
	return m.result
}

func monk (in chan(Task), out chan(int), i int) {
	fmt.Printf("Goroutine %d has started!\n", i)
	for {
		t := <- in
		fmt.Printf("Goroutine %d received task %d\n", i, t.TaskNum)
		result := Process(t.X, t.Y, t.C)
		out <- result
	}
}

func dispatcher (channel chan(Task), nr int) {
	
	for i:=0; i<nr; i++ {
		var c Component
		if i < nr/2 {
			c = new(Sum)
		} else {
			c = new(Multiply)
		}
		channel <- Task{&c, i, i+1, i}
	}
	
}

func listener (channel chan(int), num_tasks int) {
	for i:=0; i<num_tasks; i++ {
		result := <- channel
		fmt.Printf("%d. Received result %d\n", i, result) 
	}
}

func main () {
	fmt.Println("Hello!")
	in_channel := make(chan(Task))
	out_channel := make(chan(int))

	num_tasks := 1000

	for i:=0; i<10; i++ {
		go monk(in_channel, out_channel, i) 
	}

	go listener(out_channel, num_tasks)
	go dispatcher(in_channel, num_tasks)
	
	time.Sleep(10*time.Second)
	fmt.Println("Bye!")
}