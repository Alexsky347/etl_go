package services

import (
	"encoding/csv"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"strings"
)


func CreateProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	topic := KafkaTopic

	// Open the CSV file
	file, err := os.Open(CSVFilePath)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	r := csv.NewReader(file)

	// Read all records
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	if OutputModeBatch {
		// Convert all records to a single string -> ok for small files
		var sb strings.Builder
		for _, record := range records {
			sb.WriteString(fmt.Sprintf("%v\n", record))
		}
		message := []byte(sb.String())

		// Produce the entire file content as a single message
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          message,
		}, nil)
	} else {
		// Produce each record as a message -> line by line
		for _, record := range records {
			message := []byte(fmt.Sprintf("%v", record))
			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          message,
			}, nil)

			if err != nil {
				fmt.Println("Error producing message:", err)
				return
			}
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000) // Wait up to 15 seconds for message delivery
}
