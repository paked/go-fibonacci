package main

import (
	"flag"
	"log"
	"time"
)

//Done var channel (used for ending the benchmark when timer has finished)
var done = make(chan bool)

//The array of fibonacci numbers
var fibs = []int{}

//Self docced ;-D
var timer = flag.Int("time", 10, "How long you want the process to last for (default 10 seconds)")

func main() {
	//Parse flags
	flag.Parse()
	//Setup a timer to tell us when the benchmark should end
	ticker := time.NewTicker(time.Duration(*timer) * time.Second)
	//Start a generator in a seperate goroutine
	go generator()
	//Check if timer has finished
	for {
		select {
		//Timer has finished
		case <-ticker.C:
			log.Printf("%v", len(fibs))
			//Send "done = true" to the generator routine
			done <- true
			return
		}
	}
}

// Generator does the actual fibonacci busiuess
func generator() {
	//Set two starting numbers(a and b)
	a, b := 0, 1
	//Num is the end result
	num := 0
	//Bool used to monitor the for loop
	finished := false
	for !finished {
		//Fibonacci it up right under here
		// 2 + 3 = 5
		// a + b = num
		num = a + b
		b = a
		a = num

		//Add the result to the end of the fibonacci slice
		fibs = append(fibs, num)

		//Poll the done value from the main routine
		select {
		case <-done:
			//End the for loop
			finished = true
		default:
			//This somehow stops the blocking
		}
	}
}
