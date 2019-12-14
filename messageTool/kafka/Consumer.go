package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

type Consumer struct {
}

func NewConsumer() *Consumer {
	return new(Consumer)
}

func (s *Consumer) GetTopicMsgs(brokers []string, topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Printf("Consumer.CousumerMsg NewConsumer err:%s", err)
		return
	}
	//get all partitionLists from topic
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		log.Printf("Consumer.CousumerMsg Partitions err:%s ", err)
		return
	}
	for partition := range partitionList {
		offset := sarama.OffsetNewest
		cousumerPartition, err := consumer.ConsumePartition(topic, int32(partition), offset)
		if err != nil {
			log.Printf("Consumer.CousumerMsg ConsumePartition:%d:,err:%s\n", partition, err)
			return
		}
		go s.CousumerMsg(cousumerPartition)
	}
}

func (s *Consumer) CousumerMsg(cousumerPartition sarama.PartitionConsumer) {
	defer cousumerPartition.Close()
	for {
		select {
		case msg := <-cousumerPartition.Messages():
			log.Println(string(msg.Value), msg.Partition, msg.Offset)
		case err := <-cousumerPartition.Errors():
			log.Println(err)
		}
	}
}
