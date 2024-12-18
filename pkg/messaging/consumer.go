package messaging

import (
	"fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func CreateConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
		"group.id": KafkaGroupId,
		"auto.offset.reset": "earliest",
	   })
	   if err != nil {
		panic(err)
	   }
	   defer c.Close()

	   topic := KafkaTopic
	   c.SubscribeTopics([]string{topic}, nil)

	// Consume messages
	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error decoding message: %v\n", err)
			continue
		}
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}