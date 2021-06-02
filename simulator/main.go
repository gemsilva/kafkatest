package main

import (
	"fmt"
	kafka2 "github.com/gemsilva/kafkatest/application/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	kafka.Publish("olá","readtest", producer)

	for {
		_ = 1
	}
	// consumer := kafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()
	// for msg := range msgChan {
	// 	fmt.Println(string(msg.Value))
	// 	go kafka2.Produce(msg)
	// }
}
