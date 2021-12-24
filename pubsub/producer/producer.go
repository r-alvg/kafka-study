package producer

import "kafka-example/pubsub"

type Producer struct {
	//in *chan int
}
func (p Producer) Produce(send func(topic pubsub.Topic), topic pubsub.Topic) {
	p.SendToKafka(send,topic)
}

func (p Producer)SendToKafka(send func(topic pubsub.Topic),topic pubsub.Topic){
	go send(topic)
}

func NewProducer() Producer {
	return Producer{}
}
