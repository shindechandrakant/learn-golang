package main

import (
	"fmt"
	data_structure "go-fiber/data-structure"
	"sync"
	"time"
)

func main() {

	var stack = data_structure.Stack{}
	fmt.Println(stack.IsEmpty())
	stack.Push(10)
	fmt.Println(stack.Front())
	fmt.Println(stack.Length())
	stack.Push(20)
	fmt.Println(stack.Length())
	stack.Pop()
	fmt.Println(stack.Front())
	stack.Pop()
	fmt.Println(stack.Front())

	/*
		fmt.Println("Gello")
		myChannel := make(chan int, 3)
		wg := &sync.WaitGroup{}

		wg.Add(2)
		defer wg.Wait()

		go func(ch chan int, wg *sync.WaitGroup) {
			fmt.Println(<-myChannel)
			fmt.Println(<-myChannel)
			fmt.Println(<-myChannel)
			wg.Done()
		}(myChannel, wg)

		go func(ch chan int, wg *sync.WaitGroup) {

			myChannel <- 5
			myChannel <- 2
			myChannel <- 3
			wg.Done()
		}(myChannel, wg)

		//myChannel <- 5
		//myChannel <- 3
		//myChannel <- 1
		//fmt.Println(<-myChannel)
		wg.Wait()
		close(myChannel)

	*/
}

var lock sync.Mutex

var messages []string

func WaitGroup() {

	var score []int
	mutx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(wg *sync.WaitGroup, mutx *sync.Mutex) {

		fmt.Println("One")
		mutx.Lock()
		score = append(score, 1)
		mutx.Unlock()
		wg.Done()
	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.Mutex) {

		fmt.Println("Two")
		mutx.Lock()
		score = append(score, 2)
		mutx.Unlock()
		wg.Done()
	}(wg, mutx)

	go func(wg *sync.WaitGroup, mutx *sync.Mutex) {

	}(wg, mutx)

	wg.Wait()
	fmt.Println(score)
	fmt.Println(len(score))

	//go greet("Hello")
	//wg.Add(1)
	//
	//go greet("World")
	//wg.Add(1)
	//wg.Wait()
	//fmt.Println(messages)
	//fmt.Println(len(messages))
}

func greet(str string) {
	//defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		lock.Lock()
		messages = append(messages, str)
		lock.Unlock()
		fmt.Printf("%d - %s\n", i, str)
	}
}
