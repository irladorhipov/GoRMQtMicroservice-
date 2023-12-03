package main

import (
	"fmt"
	rabbitMQ "irladorhipov/internal/rmq"
)

func main() {
	conn, err := rabbitMQ.Connect()
	if err != nil {
		fmt.Println(conn)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("[*] - waiting for messages")
	<-forever
}
