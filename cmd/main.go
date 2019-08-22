package main

import (
	"github.com/restaurant/pkg/database/mysql"
	"github.com/restaurant/pkg/server"
)

func main() {
	db := mysql.NewMySqlDB("restaurant_database_1","hcs123","hcs12345","3306","restaurants")
	defer db.Close()
	s := server.NewServer(db)
	s.Start(":4000")
}
