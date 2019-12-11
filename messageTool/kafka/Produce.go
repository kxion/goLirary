package kafka

import (
	"github.com/Shopify/sarama"
	"log"
)

type Produce struct {
}

func NewProduce() *Produce {
	return new(Produce)
}

func (s *Produce) SendSync(brokers []string, topic, key string, value []byte) {
	cfg := sarama.NewConfig()
	cfg.Producer.Compression = sarama.CompressionGZIP
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.Return.Errors = true
	cfg.Producer.MaxMessageBytes = 3000000
	cfg.Producer.Timeout = 30
	cfg.Producer.Return.Successes = true
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
	}
	msg.Value = sarama.ByteEncoder(value)
	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		log.Printf("LangYanKfkSender.SendMsg NewSyncProducer err:%s", err.Error())
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("LangYanKfkSender.SendMsg producer.close err:%s", err.Error())
		}
	}()
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("LangYanKfkSender.SendMsg send message err:%s", err.Error())
	} else {
		log.Printf("message sent to partition %d at offset %d\n", partition, offset)
	}
}

func (s *Produce) SendAsync(brokers []string, topic, key string, value []byte) {
	//cfg := sarama.NewConfig()
	//cfg.Producer.Compression = sarama.CompressionGZIP
	//cfg.Producer.RequiredAcks = sarama.WaitForLocal
	//cfg.Producer.Partitioner = sarama.NewHashPartitioner
	//cfg.Producer.Return.Errors = true
	//cfg.Producer.MaxMessageBytes = 3000000
	//cfg.Producer.Timeout = 30
	//cfg.Producer.Return.Successes = true
	//msg := &sarama.ProducerMessage{
	//	Topic: topic,
	//	Key:   sarama.StringEncoder(key),
	//}
	//msg.Value = sarama.ByteEncoder(value)
	//producer, err := sarama.NewAsyncProducer(brokers, cfg)
	//if err != nil {
	//	log.Printf("LangYanKfkSender.SendMsg NewAsyncProducer err:%s", err.Error())
	//}

}
