package main

import (
	"github.com/restaurant/Res_Man_MicroService/pkg/database/mysql"
	"github.com/restaurant/Res_Man_MicroService/pkg/server"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	db := mysql.NewMySqlDB("localhost", "root", "password", "3306", "Restaurant")
	defer db.Close()
	err := mysql.MigrateDatabase(db)
	if err != nil {
		panic("can't migrate db: " + err.Error())
	}
	s := server.NewServer(db)
	println(port)
	s.Start(":" + "4000")
}
