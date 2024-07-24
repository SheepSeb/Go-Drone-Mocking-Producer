# Go Mocking Data Producer
This project creates a mock data producer using Go and Kafka. The data producer will generate random data and send it to a Kafka topic.
## How to start the Kafka Cluster using Docker with docker-compose

Using the ```docker-compose.yml``` file we can start the Kafka cluster with Zookeeper and Kafka in separate containers. We have 4 services in the cluster:
- Zookeeper
- Kafka (2 brokers)
- KafkaUI

```bash
docker-compose up -d
```

## How to start the Kafka Cluster using Docker with separate docker containers

Note: This might not work as expected and you should try to use the docker-compose file.

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
  -e KAFKA_CFG_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
  -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT \
  -e KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://kafka:29092
  -e KAFKA_LISTENERS: PLAINTEXT://0.0.0.0:9092,PLAINTEXT_INTERNAL://0.0.0.0:29092 \
  -e KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT_INTERNAL \
  -e ALLOW_PLAINTEXT_LISTENER=yes \
  bitnami/kafka:latest
```

Optional for UI add KafkaUI
```bash
docker run -d \
  --name kafka-ui \
  --network kafka-network \
  -p 8081:8080 \
  -e KAFKA_CLUSTERS_0_NAME=local \
  -e KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=kafka1:29092,kafka2:29093 \
  -e KAFKA_CLUSTERS_0_ZOOKEEPER=zookeeper:2181 \
  provectuslabs/kafka-ui:v0.6.2
```

Access the UI at http://localhost:8081, we chose to start on the 8081 port due to the Java Spring App on the 8080 port.