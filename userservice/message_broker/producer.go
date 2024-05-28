package message_broker

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IBM/sarama"

	"userservice/schemas"
)

var producer sarama.SyncProducer
var topicName string

func InitMessageProducer() {
	topicName = os.Getenv("KAFKA_TOPIC")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	broker := os.Getenv("KAFKA_BROKER")
	var err error
	producer, err = sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatal(err)
	}
}

func SendEventToBroker(event schemas.Event) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Failed to marshal event: %v", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topicName,
		Partition: 0,
		Value:     sarama.ByteEncoder(eventBytes),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	return nil
}
