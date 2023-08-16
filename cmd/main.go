package main

import (
	"fmt"
	"time"
)

type Subscriber chan string

func main() {
	publisher := make(Subscriber)
	subscriber1 := make(Subscriber)
	subscriber2 := make(Subscriber)

	go publish(publisher)
	go subscribe(subscriber1, "Subscriber 1")
	go subscribe(subscriber2, "Subscriber 2")

	for {
		select {
		case event := <-publisher:
			subscriber1 <- event
			subscriber2 <- event
		}
	}
}

func publish(publisher Subscriber) {
	for {
		time.Sleep(time.Second)
		publisher <- "New event"
	}
}

func subscribe(sub Subscriber, name string) {
	for {
		event := <-sub
		fmt.Printf("%s received: %s\n", name, event)
	}
}
