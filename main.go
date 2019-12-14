package main

import (
	"goLirary/messageTool/kafka"
)

func main() {
	openChan := make(chan int)
	//produce := kafka.NewProduce()
	cousumer := kafka.NewConsumer()
	brokers := []string{"127.0.0.1:9092"}
	topic := "xiaoheiwa"
	//go func() {
	//	for i := 0; i < 40; i++ {
	//		produce.SendSync(brokers, topic, "ceshi", []byte(fmt.Sprintf("test:%d", i)))
	//		time.Sleep(5 * time.Second)
	//	}
	//}()
	go cousumer.GetTopicMsgs(brokers, topic)
	<-openChan
}
