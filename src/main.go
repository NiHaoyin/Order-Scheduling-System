package main

import (
	"fmt"
	"time"
)

type MessageQueue struct {
	queue chan string
}

func NewMessageQueue(size int) *MessageQueue {
	return &MessageQueue{
		queue: make(chan string, size),
	}
}

func (mq *MessageQueue) Enqueue(message string) {
	mq.queue <- message
}

func (mq *MessageQueue) Dequeue() string {
	return <-mq.queue
}

func main() {
	messageQueue := NewMessageQueue(3)

	// Producer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			message := fmt.Sprintf("Message %d", i)
			messageQueue.Enqueue(message)
			fmt.Printf("Produced: %s\n", message)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Consumer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			message := messageQueue.Dequeue()
			fmt.Printf("Consumed: %s\n", message)
			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for goroutines to finish
	time.Sleep(10 * time.Second)
}
