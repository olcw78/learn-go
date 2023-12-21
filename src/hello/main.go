package main

import (
	"fmt"
	"time"
)

func main() {
	// say("Sync")

	// go say("Async1")
	// go say("Async2")
	// go say("Async3")

	// time.Sleep(time.Second * 3)

	// WaitGroup 생성, 2 개의 Goroutine 대기
	// var wait sync.WaitGroup
	// wait.Add(2)

	// go func() {
	// 	defer wait.Done()
	// 	fmt.Println("Hello")
	// }()

	// go func(msg string) {
	// 	defer wait.Done()
	// 	fmt.Println(msg)
	// }("Hi")

	// wait.Wait() // goroutine 모두 끝날때까지 대기

	// 4개의 CPU 사용
	// runtime.GOMAXPROCS(4)

	// ch := make(chan int)

	// go func() {
	// 	ch <- 123
	// }()

	// i := <-ch

	// println(i)

	// done := make(chan bool)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(i)
	// 	}
	// 	done <- true
	// }()

	// <-done

	// ch := make(chan int)
	// ch <- 1 // 수신 루틴이 없어서 데드락!
	// fmt.Println(<-ch)

	// ch := make(chan int, 1)

	// ch <- 101

	// fmt.Println(<-ch)

	// ch := make(chan string, 1)

	// send(ch)
	// fmt.Println(<-ch)
	// ch <- "1000"
	// recv(ch)

	// ch := make(chan int, 2)

	// ch <- 1
	// ch <- 2

	// close(ch)

	// for {
	// 	if i, success := <-ch; success {
	// 		println(i)
	// 	} else {
	// 		break
	// 	}
	// }

	// for i := range ch {
	// 	println(i)
	// }

	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 done")
		case <-done2:
			println("run2 done")
			break EXIT
		}
	}
}

func run1(done chan<- bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan<- bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func send(ch chan<- string) {
	ch <- "Data"
	// x := <-ch
}

func recv(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}
