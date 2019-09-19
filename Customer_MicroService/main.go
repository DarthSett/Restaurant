package main

import (
	"github.com/restaurant/Customer_MicroService/pkg/database/mysql"
	"github.com/restaurant/Customer_MicroService/pkg/server"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	db := mysql.NewMySqlDB("localhost", "root", "password", "3306", "restaurant")
	defer db.Close()
	s := server.NewServer(db)
	println(port)
	s.Start(":" + "4000")
}
