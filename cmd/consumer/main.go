package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/mauFade/go-intensive/internal/infra/database"
	"github.com/mauFade/go-intensive/internal/usecase"
	"github.com/mauFade/go-intensive/pkg/kafkaclient"
)

func main() {
	conn, err := sql.Open("sqlite3", "./orders.bd")

	if err != nil {
		panic(err)
	}

	defer conn.Close() // Defer -> fecha o banco depois que tudo for executado

	repository := database.NewOrderRepository(conn)

	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	kafkaMessageChannel := make(chan *kafka.Message)

	topics := []string{"orders"}
	host := "host.docker.internal:9094"

	go kafkaclient.Consume(topics, host, kafkaMessageChannel)

	kafkaWorker(kafkaMessageChannel, usecase)
}

func kafkaWorker(messageChan *kafka.Message, uc usecase.CalculateFinalPrice) {
	for msg := range messageChan {
		var orderInputDTO usecase.OrderInputDTO

		err := json.Unmarshal(msg.Value, &orderInputDTO)

		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(orderInputDTO)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Kafka has processed the order: %s\n", outputDTO.ID)
	}
}
