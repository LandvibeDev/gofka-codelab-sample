package kafka

import (
	"fmt"
	"github.com/LandvibeDev/gofka-codelab-sample/config"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	"os/signal"
	"syscall"
)

func NewConsumerConnector(kafkaConfig config.KafkaConfiguration) (*ConsumerConnector, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               kafkaConfig.Hosts,
		"group.id":                        kafkaConfig.Consumer.GroupId,
		"session.timeout.ms":              kafkaConfig.Consumer.SessionTimeoutMs,
		"auto.offset.reset":               kafkaConfig.Consumer.AutoOffsetReset,
		"enable.partition.eof":            true,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
	})

	if err != nil {
		return nil, err
	}
	return &ConsumerConnector{
		consumer: c,
		config:   kafkaConfig,
	}, nil
}

type ConsumerConnector struct {
	consumer *kafka.Consumer
	config   config.KafkaConfiguration
}

func (c *ConsumerConnector) StartPeek() error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	err := c.consumer.Subscribe(c.config.Consumer.Topic, nil)
	if err != nil {
		return err
	}
	defer c.consumer.Close()

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-c.consumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.consumer.Unassign()
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	return nil
}
