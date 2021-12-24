package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

func GenerateRandomFloat() string {
	return fmt.Sprintf("%.2f", rand.Float64()*1000)
}
func RandomTID() string {
	id := uuid.New()
	return id.String()
}
func RandomCID() string {
	return fmt.Sprintf("%s-%d", "02", rand.Intn(999999-100000))
}
func RandomIntallment() string {
	return fmt.Sprintf("%d", rand.Intn(12-1))
}
func InterfaceMarshal(payload interface{}) []byte {
	response, err := json.Marshal(payload)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
	return response
}

func GeneratedCardPayloadData(quantity int) []Payload {
	var jobs []card.Payload
	for i := 0; i < quantity; i++ {
		jobs = append(jobs, Payload{
			TypePgt:     "credit card",
			Value:       GenerateRandomFloat(),
			CustomerID:  RandomCID(),
			Tid:         RandomTID(),
			Installment: RandomIntallment(),
			Count:       strconv.Itoa(i),
		})
	}
	return jobs
}
