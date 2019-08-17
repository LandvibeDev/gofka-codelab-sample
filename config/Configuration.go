package config

import (
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server  ServerConfiguration
	MongoDb DatabaseConfiguration
	Kafka   KafkaConfiguration
}

type DatabaseConfiguration struct {
	Hosts string
}

type ServerConfiguration struct {
	Port int
}

type KafkaConfiguration struct {
	Hosts    string
	Consumer KafkaConsumerConfiguration
	Topic    KafkaTopicConfiguration
}

type KafkaConsumerConfiguration struct {
	Topic            string `mapstructure:"topic"`
	GroupId          string `mapstructure:"groupid"`
	SessionTimeoutMs string `mapstructure:"session-timeout-ms"`
	AutoOffsetReset  string `mapstructure:"auto-offset-reset"`
}

type KafkaTopicConfiguration struct {
	Name              string `mapstructure:"name"`
	NumPartitions     int    `mapstructure:"num-partitions"`
	ReplicationFactor int    `mapstructure:"replication-factor"`
}

func LoadConfiguration(log echo.Logger) Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}
