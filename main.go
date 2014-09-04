package main

import (
	"log"
	"time"
)

var done = make(chan bool)
var fibs = make(chan []int)

func main() {
	ticker := time.NewTicker(10 * time.Second)
	go generator()
	for {
		select {
		case <-ticker.C:
			log.Println("DONE")
			done <- true
			return
		}
	}
}

func generator() {
	a, b := 0, 1
	num := 0
	finished := false
	for {
		num = a + b
		b = a
		a = num
		fibs = append(fibs, num)
		finished = <-done
		if finished {
			return
		}
	}
}
