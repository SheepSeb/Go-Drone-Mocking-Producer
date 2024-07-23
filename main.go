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
		"client.id":         "goapp",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	topic_str := "drone_topic"

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic_str, Partition: kafka.PartitionAny},
		Value:          drone.ToJson()},
		nil, // delivery channel
	)
}
