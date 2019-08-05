package main

import (
	"github.com/restaurant/pkg/database/mysql"
	"github.com/restaurant/pkg/server"
)

func main() {
	db := mysql.NewMySqlDB("127.0.0.1","root","Zamorak1","3306","Restaurant")
	defer db.Close()
	s := server.NewServer(db)
	s.Start(":3000")
}
