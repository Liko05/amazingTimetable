package main

import (
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

var messages = make(chan string)
var isFinished = make(chan bool)

func main() {
	//print the number of processors
	println("Number of processors: " + strconv.Itoa(runtime.NumCPU()))
	// print available threads
	println("Number of threads: " + strconv.Itoa(runtime.GOMAXPROCS(0)))

	go produceMessage()
	go consumeMessage()

	<-isFinished
}

func hardComputation() int {
	timeStart := time.Now()
	randNum := rand.Intn(1000000000-100000) + 100000
	for i := 0; i < randNum; i++ {
		math.Pow(2, 2)
	}
	return int(time.Since(timeStart).Milliseconds())
}

func produceMessage() {
	timeStart := time.Now()
	// run 16 threads to produce messages for 100 seconds and then stop
	for i := 0; i < 11; i++ {
		go func() {
			for {
				select {
				case <-time.After(100 * time.Millisecond):
					timeStamp := time.Now().String()
					messages <- "Time: " + timeStamp + " Number: " + strconv.Itoa(hardComputation())
				}
				if time.Since(timeStart) > 100*time.Second {
					isFinished <- true
					return
				}
			}
		}()
	}
}

func consumeMessage() {
	for i := 0; i < 4; i++ {
		go func(id int) {
			for {
				select {
				case message := <-messages:
					timeStamp := time.Now().String()
					println("Thread: " + strconv.Itoa(id) + "Time: " + timeStamp + " Message: " + message)
				}
			}
		}(i)
	}

	<-isFinished
}
