package main

import (
	"bytes"
	"drone_producer/domain"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io"
	"net/http"
	"os"
	"time"
)

func simpleHttps() {
	var drone domain.Drone
	drone = domain.NewDrone(1, "seb", false)
	println(drone.String())

	jsonPayload, err := json.Marshal(drone)

	if err != nil {
		println(err)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/drone", bytes.NewBuffer(jsonPayload))

	if err != nil {
		println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)
}

func main() {

	var drone domain.Drone
	drone = domain.NewDrone(1, "seb", false)
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "go app",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer p.Close()

	topicStr := "drone_topic"
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, err := json.Marshal(drone)
			if err != nil {
				fmt.Printf("Failed to marshal drone: %s\n", err)
				continue
			}
			err = p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topicStr, Partition: kafka.PartitionAny},
				Value:          drone.ToJson()},
				nil, // delivery channel
			)

			if err != nil {
				fmt.Printf("Failed to produce message: %s\n", err)
			} else {
				fmt.Printf("Produced message: \n")
			}

		}
	}
}
