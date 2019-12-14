package main

import (
	"goLirary/messageTool/kafka"
	"time"
)

func main() {
	produce := kafka.NewProduce()
	cousumer := kafka.NewConsumer()
	brokers := []string{"127.0.0.1:9092"}
	topic := "xiaoheiwa"
	go produce.SendMultiMsg(brokers, topic, "xiaohiewa")
	go cousumer.GetTopicMsgs(brokers, topic)
	time.Sleep(10 * time.Minute)
}
