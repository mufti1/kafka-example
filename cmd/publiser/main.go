package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/mufti1/kafka-example/producer"
	"github.com/sirupsen/logrus"
)

func main() {
	// Setup Logging
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	kafkaConfig := getKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return
	}
	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync-producer")

	kafka := &producer.KafkaProducer{
		Producer: producers,
	}

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("message number %v", i)
		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			panic(err)
		}
	}
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
