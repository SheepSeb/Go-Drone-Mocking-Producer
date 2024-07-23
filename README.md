# Go Mocking Data Producer
This project creates a mock data producer using Go and Kafka. The data producer will generate random data and send it to a Kafka topic.
## How to start the Kafka Cluster using Docker

Create the network for the Kafka cluster
```bash
docker network create kafka-network
```

Start the Zookeeper
```bash
docker run -d --name zookeeper --network kafka-network -p 2181:2181 -e ALLOW_ANONYMOUS_LOGIN=yes bitnami/zookeeper:latest
```

Start the Kafka container and expose ports
```bash
docker run -d --name kafka --network kafka-network \
  -p 9092:9092 \
  -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 \
  -e ALLOW_PLAINTEXT_LISTENER=yes \
  bitnami/kafka:latest
```

Optional for UI add KafkaHQ
```bash
docker run -d --name akhq --network kafka-network \
  -p 8081:8080 \
  -e AKHQ_CONFIGURATION='{
    "akhq": {
      "connections": {
        "my-cluster": {
          "properties": {
            "bootstrap.servers": "kafka:9092"
          }
        }
      }
    }
  }' \
  tchiotludo/akhq
```

Access the UI at http://localhost:8081, we chose to start on the 8081 port due to the Java Spring App on the 8080 port.