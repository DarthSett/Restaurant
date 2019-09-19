package main

import (

	"github.com/restaurant/Reservation_MicroService/pkg/database/mysql"
	"github.com/restaurant/Reservation_MicroService/pkg/server"
	"github.com/streadway/amqp"
//	"log"
	"os"
//	"strconv"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	db := mysql.NewMySqlDB("localhost", "root", "password", "3306", "reservation")
	defer db.Close()
	println("calling rconnect")
	conn := rConnect()
	println("leaving rconnect")
	defer conn.Close()



	s := server.NewServer(db)
	println(port)
	s.Start(":" + "4000")
}


func rConnect() *amqp.Connection {

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		println(err.Error())
		println("trying to rc")
		time.Sleep(5 * time.Second)
		return rConnect()
	}

	return conn
}