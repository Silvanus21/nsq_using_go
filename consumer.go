package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"sync"
)

func consumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("first_topic", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
