package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type Produce struct {
}

func NewProduce() *Produce {
	return new(Produce)
}

func (s *Produce) GetSyncProducer(brokers []string) sarama.SyncProducer {
	cfg := sarama.NewConfig()
	cfg.Producer.Compression = sarama.CompressionGZIP
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.Return.Errors = true
	cfg.Producer.MaxMessageBytes = 3000000
	cfg.Producer.Timeout = 30
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		log.Printf("Produce.SendSync err:%s", err.Error())
	}
	return producer
}

func (s *Produce) SendSyncMsg(producer sarama.SyncProducer, brokers []string, topic, key string, value []byte) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
	}
	msg.Value = sarama.ByteEncoder(value)
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Produce.SendSync send message err:%s", err.Error())
	} else {
		log.Printf("Produce.SendSync message sent msg:%s", string(value))
	}
}

func (s *Produce) GetAsyncProducer(brokers []string) sarama.AsyncProducer {
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Compression = sarama.CompressionSnappy
	cfg.Producer.Flush.Frequency = 500 * time.Millisecond
	producer, err := sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		log.Printf("Produce.SendAsync  err:%s", err.Error())
	}
	return producer
}

func (s *Produce) SendAsyncMsg(producer sarama.AsyncProducer, brokers []string, topic, key string, value []byte) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
	}
	msg.Value = sarama.ByteEncoder(value)
	producer.Input() <- msg
	go func() {
		for range producer.Successes() {
			log.Printf("Produce.SendSync message sent msg:%s", string(value))
		}
	}()
}

func (s *Produce) SendMultiMsg(brokers []string, topic, key string) {
	producer := s.GetSyncProducer(brokers)
	for i := 0; i < 40; i++ {
		go s.SendSyncMsg(producer, brokers, topic, "ceshi", []byte(fmt.Sprintf("test:%d", i)))
		time.Sleep(1 * time.Second)
	}
}
