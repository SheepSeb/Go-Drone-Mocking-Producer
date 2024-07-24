package main

import (
	"bytes"
	"drone_producer/domain"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"io"
	"net/http"
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

	const kafka_ip = "localhost:9092"
	const kafka_topic = "drone-topic"

	var drone domain.Drone
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{kafka_ip}, config)

	if err != nil {
		fmt.Println("Failed to start Sarama producer: ", err)
		return
	}

	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			fmt.Println("Failed to close Sarama producer: ", err)
		}
	}(producer)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			drone = domain.NewDrone(1, "seb", false)
			value, err := json.Marshal(drone)
			if err != nil {
				fmt.Printf("Failed to marshal drone: %s\n", err)
				continue
			}

			msg := &sarama.ProducerMessage{
				Topic: kafka_topic,
				Value: sarama.StringEncoder(value),
			}

			partition, offset, err := producer.SendMessage(msg)

			if err != nil {
				fmt.Printf("Failed to produce message: %s\n", err)
				continue
			} else {
				fmt.Printf("Produced message: partition %d, offset %d\n", partition, offset)
			}

		}
	}
}
