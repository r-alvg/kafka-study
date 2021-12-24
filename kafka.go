package main

import (
	"fmt"
	"kafka-example/db"
	"kafka-example/paymentType/card"

	"kafka-example/pubsub/consumer"
	"kafka-example/pubsub/producer"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func main() {

	db.Start()

	const nConsumers = 5
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())

	p := producer.NewProducer()
	c := consumer.NewConsumer()
	wg.Add(nConsumers)

	card.SendToKafkaTopic(&wg,p,10)


	fmt.Println("\nStart receiving from Kafka") // start consumers
	topicName := "CARD_PGT"
	go c.Work(&wg,nConsumers,topicName)



	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	wg.Wait()
}

