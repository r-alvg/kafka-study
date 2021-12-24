package pubsub

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"kafka-example/helpers"
	"sync"
)


func NewKafkaProducer() *kafka.Producer {
	server := fmt.Sprintf("%s:%s","a7a14634f555441ccbe2ca8731c63823-744881338.us-east-1.elb.amazonaws.com","19092")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		panic(err)
	}
	return p
}


func SendToCardPgtTopic(topic Topic)  {
	SendToKafka(topic.NameTP(),topic)
}

func SendToKafka(topic string, payload interface{}) {
	var wg sync.WaitGroup
	//var m sync.Mutex
	wg.Add(2)

	//currentMassage := make(chan CardPgt, 1)
	deliveryChan := make(chan kafka.Event, 10000)
	// Produce messages to topic (asynchronously)
	go func() {
		topic := topic
		sendMessage(NewKafkaProducer(),topic,deliveryChan,payload)
		wg.Done()
	}()

	// Delivery report handler for produced messages
	DeliveryReport(&wg,deliveryChan)

	wg.Wait()

}
func DeliveryReport(wg *sync.WaitGroup,deliveryChan chan kafka.Event)  {
	go func() {
		for e := range deliveryChan {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)

				} else {
					//fmt.Printf("Delivered message to topic: %s partition: %v\n", *ev.TopicPartition.Topic,ev.TopicPartition.Partition)
				}
			}
		}
		wg.Done()
	}()
}
func sendMessage(p *kafka.Producer,topic string,delivery chan kafka.Event ,payload interface{})  {
	defer p.Close()
	_ = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          helpers.InterfaceMarshal(payload),
	}, delivery)
	//SendToKafka("LOGS_PGT",[]{fmt.Sprintf("SEND: %+v to CARD_PGT topic",v)})
}

func ReceiveFromTopic(topicName string) {
	var mutex sync.Mutex
	server := fmt.Sprintf("%s:%s","a7a14634f555441ccbe2ca8731c63823-744881338.us-east-1.elb.amazonaws.com","19092")
	fmt.Println("\nStart receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "0",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	cc := 0
	_ = c.SubscribeTopics([]string{topicName}, nil)
	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			mutex.Lock()
			cc++
			mutex.Unlock()

			fmt.Printf("Received from Kafka %s: %s count: %v\n", msg.TopicPartition, string(msg.Value), cc)

		} else {
			fmt.Printf("\nConsumer error: %v (%v)\n", err, msg)
			break
		}
	}

	_ = c.Close()


}