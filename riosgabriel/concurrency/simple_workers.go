package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var w sync.WaitGroup

func populateChan(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func worker(i int, c chan int) {
	for val := range c {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond) //act like we're doing something
		fmt.Println("worker", i, "recieved", val)
	}
	fmt.Println("worker", i, "done")
	w.Done()
}

func main() {
	//define worker count
	workers := 3

	//create channel
	c := make(chan int)

	//populate the channel
	go populateChan(c)

	//tell sync how many goroutines are out there
	w.Add(workers)

	//start workers to read from channel
	for i := 1; i <= workers; i++ {
		go worker(i, c)
	}

	//wait for goroutines to complete
	w.Wait()
}