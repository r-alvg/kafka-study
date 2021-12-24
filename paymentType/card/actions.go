package card

import (
	"encoding/json"
	"fmt"
	"kafka-example/helpers"
	"kafka-example/pubsub"
	"kafka-example/pubsub/producer"
	"math/rand"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

func SendToKafkaTopic(wg *sync.WaitGroup,p producer.Producer, quantity int)  {
	fmt.Printf("start send messages to kafka")
	for _, v := range GeneratedCardPayloadData(quantity) {
		wg.Add(1)
		go p.Produce(
			pubsub.SendToCardPgtTopic,
			pubsub.CardPgtTopic{
				Name: "CARD_PGT",
				Payload: v,
			})
		wg.Done()
	}
}
