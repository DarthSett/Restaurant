package helpers

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func SendSignal(body map[int]int) {
	conn := rConnect()
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil{
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil{
		panic(err)
	}
	b,err := json.Marshal(body)
	if err != nil{
		panic(err)
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         b,
		})
	if err != nil{
		panic(err)
	}
	log.Printf(" [x] Sent %v", string(b))
}

func rConnect() *amqp.Connection {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		println(err.Error())
		println("trying to rc")
		time.Sleep(5 * time.Second)
		return rConnect()
	}

	return conn
}