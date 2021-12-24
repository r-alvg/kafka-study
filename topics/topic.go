package topics

type Topic interface {
	NameTP() string
	PayloadTP() interface{}
}

type topic struct {
	Name string
	Payload interface{}
}

func (t topic) NameTP() string {
	return t.Name
}

func (t topic) PayloadTP() interface{} {
	return t.Payload
}

func CreateTopic(name string, payload interface{}) topic  {
	return topic{Name: name , Payload: payload}
}