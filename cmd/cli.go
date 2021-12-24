package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"strconv"
	"github.com/segmentio/kafka-go"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "create topic",
				Usage: "<broker> <topic> <partition-count> <replication-factor>",
				Action: func(c *cli.Context) error {

					broker := os.Args[1]
					topic := os.Args[2]
					numParts, err := strconv.Atoi(os.Args[3])
					if err != nil {
						fmt.Printf("Invalid partition count: %s: %v\n", os.Args[3], err)
						os.Exit(1)
					}
					replicationFactor, err := strconv.Atoi(os.Args[4])
					if err != nil {
						fmt.Printf("Invalid replication factor: %s: %v\n", os.Args[4], err)
						os.Exit(1)
					}

					partition := 0
					conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, partition)
					if err != nil {
						panic(err.Error())
					}
					defer conn.Close()
					topicConfigs := []kafka.TopicConfig{
						kafka.TopicConfig{
							Topic:             topic,
							NumPartitions:     numParts,
							ReplicationFactor: replicationFactor,
						},
					}

					err = conn.CreateTopics(topicConfigs...)
					if err != nil {
						panic(err.Error())
					}

					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)

	if err != nil{
		log.Printf("Erro ao executar comando - %s", err)
		return
	}
}
