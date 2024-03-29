package kafkaclient

import ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func Consume(topics []string, server string, msgChannel chan *ckafka.Message) {
	consumer, err := ckafka.NewConsumer(ckafka.ConfigMap{
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
