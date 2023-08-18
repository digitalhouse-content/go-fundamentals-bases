package main

import (
	"fmt"
	"time"
)

func main() {

	go myProcess("A")
	go myProcess("B")
	time.Sleep(3 * time.Second)

	myFirstChannel := make(chan string)

	go myProcessWithChannel("C", myFirstChannel)

	result := <-myFirstChannel
	fmt.Println(result)
	close(myFirstChannel)

	channelA := make(chan string)
	channelB := make(chan string)

	go myProcessWithChannel("D", channelA)
	go myOtherProcessWithChannel("E", channelB)

	resultA := <-channelA
	fmt.Println("A: ", resultA)

	resultB := <-channelB
	fmt.Println("B: ", resultB)

	close(channelA)
	close(channelB)
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 5 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	c <- "ok"
}

func myOtherProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 20 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}

	c <- "ok"
}

func myProcess(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}
