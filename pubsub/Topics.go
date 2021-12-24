package pubsub

import "kafka-example/paymentType/card"

type Topic interface {
	NameTP() string
	PayloadTP() interface{}
}


type CardPgtTopic struct {
	Name string
	Payload card.Payload
}
func (c CardPgtTopic) NameTP() string {
	return c.Name
}
func (c CardPgtTopic) PayloadTP() interface{} {
	return c.Payload
}



