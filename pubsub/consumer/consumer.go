package consumer

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"sync"
)

type Consumer struct {
	nConsumers int
}
func (c Consumer) Work(wg *sync.WaitGroup,nConsumers int,topicName string) {
	defer wg.Done()
	fmt.Printf("starting consumers")
	for i :=0 ; i <= nConsumers ; i++{
		go c.kafkaConsumer(topicName,i)
	}
}
//func (c Consumer) Consume(ctx context.Context) {
//	for {
//		select {
//		case job := <-*c.in:
//			c.nConsumers <- job
//		case <-ctx.Done():
//			close(c.nConsumers)
//			return
//		}
//	}
//}

func NewConsumer() Consumer {
	return Consumer{}
}

var Count = 0
func (c Consumer) kafkaConsumer(topicName string,consumerId int)  {
	var mutex sync.Mutex
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "xxxxxxxxxxxxxxx.us-east-1.elb.amazonaws.com:19092",
		"group.id":          "card-send",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	_ = consumer.SubscribeTopics([]string{topicName}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			mutex.Lock()
			Count++
			fmt.Printf("consumer:%v Received from Kafka %s: %s count: %v\n", consumerId,msg.TopicPartition, string(msg.Value), Count)
			mutex.Unlock()

		} else {
			fmt.Printf("\nConsumer error: %v (%v)\n", err, msg)
			break
		}
	}
	_ = consumer.Close()
}
