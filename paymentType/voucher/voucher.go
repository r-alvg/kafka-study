package voucher

import "encoding/json"

type Payload struct {
	TypePgt 	string 	`json:"typePgt" bson:"typePgt"`
	Value   	string	`json:"value" bson:"value"`
	Tid     	string  `json:"tid" bson:"tid"`
	CustomerID	string  `json:"customerID" bson:"customerID"`
	Count		string  `json:"count" bson:"count"`
}

func (c Payload) ToString()  string {
	Job,_ := json.Marshal(c)
	return string(Job)
}

type Topic struct {
	Name string
	Payload Payload
}

func (c Topic) NameTP() string {
	return c.Name
}
func (c Topic) PayloadTP() Payload{
	return c.Payload
}
