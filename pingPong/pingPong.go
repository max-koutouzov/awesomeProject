package main

import (
	"fmt"
	"time"
)

func main() {
	ping, pong := make(chan int), make(chan int)
	var counter int64

	t := time.NewTimer(1 * time.Minute)
	done := make(chan struct{})
	shutdown := make(chan struct{})

	// ping->pong
	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-ping:
				counter++
				pong <- v
			}
		}
		done <- struct{}{}
	}()

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-pong:
				ping <- v
			}
		}
		done <- struct{}{}
	}()

	ping <- 1

	<-t.C
	close(shutdown)

	select {
	case <-ping:
	case <-pong:
	}

	<-done
	<-done
	t.Stop()
	fmt.Println(counter)
}
