package main

import (
   "github.com/nsqio/go-nsq"
   "log"
   "strconv"
)

func producer() {
   config := nsq.NewConfig()
   w, _ := nsq.NewProducer("127.0.0.1:4150", config)
   for i:= 0; i < 1000; i++ {
      err := w.Publish("first_topic", []byte("test" + strconv.Itoa(i)))
      if err != nil {
         log.Panic("Could not connect")
      }
   }
   w.Stop()
}