package message_broker

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IBM/sarama"

	"statisticsservice/schemas"
)

func processMessage(msg *sarama.ConsumerMessage) {
	if msg == nil {
		return
	}
	var event schemas.Event
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		log.Printf("Failed to unmarshal event: %v\n", err)
	}
	log.Println(event)
}

func RunConsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	broker := os.Getenv("KAFKA_BROKER")
	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		log.Fatal(err)
	}

	topic := os.Getenv("KAFKA_TOPIC")
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case msg, ok := <-partitionConsumer.Messages():
				if !ok {
					consumer.Close()
					partitionConsumer.Close()
					break
				} else {
					processMessage(msg)
				}
			case err := <-partitionConsumer.Errors():
				log.Printf("Error consuming message: %v\n", err)
			}
		}
	}()
}
