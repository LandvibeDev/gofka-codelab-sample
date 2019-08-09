package kafka

import (
	"context"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"time"
)

type KafkaConfig struct {
	Hosts string
}

type KafkaConsumerConfig struct {
	KafkaConfig
	Topic            string
	GroupId          string
	MaxPeek          int
	SessionTimeoutMs time.Duration
	AutoOffsetReset  kafka.Offset
}

type KafkaTopicConfig struct {
	Topic             string
	NumPartitions     int
	ReplicationFactor int
}

func GetProducer(kafkaConfig KafkaConfig) (*ProducerConnector, error) {
	return NewProducerConnector(kafkaConfig)
}

func EnsureTopic(topicConfig KafkaTopicConfig, kafkaConfig KafkaConfig) ([]kafka.TopicResult, error) {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": kafkaConfig.Hosts})
	if err != nil {
		return nil, err
	}
	defer admin.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout, err := time.ParseDuration("60s")
	if err != nil {
		return nil, err
	}

	results, err := admin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:             topicConfig.Topic,
			NumPartitions:     topicConfig.NumPartitions,
			ReplicationFactor: topicConfig.ReplicationFactor}},
		kafka.SetAdminOperationTimeout(timeout),
	)

	if err != nil {
		fmt.Printf("Failed to create topic: $d\n", topicConfig.Topic, err)
		return nil, err
	}

	return results, nil
}
