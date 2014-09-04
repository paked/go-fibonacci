package main

import (
	"log"
	"time"
)

var done = make(chan bool)
var fibs = []int{}

func main() {
	ticker := time.NewTicker(10 * time.Second)

	go generator()
	for {
		select {
		case <-ticker.C:
			log.Println("DONE")
			log.Printf("%v", len(fibs))
			done <- true
			return
		}
	}
}

func generator() {
	a, b := 0, 1
	num := 0
	finished := false
	for !finished {
		num = a + b
		b = a
		a = num
		fibs = append(fibs, num)

		select {
		case <-done:
			finished = true
		default:
		}
	}
}
