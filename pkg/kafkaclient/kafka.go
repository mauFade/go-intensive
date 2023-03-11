package kafkaclient

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func Consume(topics []string, server string, msgChannel chan *kafka.Message) {
	consumer, err := kafka.NewConsumer(kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "goapp",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			msgChannel <- msg
		}
	}
}
