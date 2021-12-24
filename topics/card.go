package topics

import "kafka-example/paymentType/card"

type CardPgtTopic struct {
	Name string
	Payload []card.Payload
}

func (c CardPgtTopic) NameTP() string {
	return c.Name
}
func (c CardPgtTopic) PayloadTP() []card.Payload{
	//var pl []interface{}
	//for _, v := range c.Payload {
	//	pl = append(pl,v)
	//}
	return c.Payload
}