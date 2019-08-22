package main

import (
	"github.com/restaurant/pkg/database/mysql"
	"github.com/restaurant/pkg/server"
	"os"
)

func main() {
	port:=os.Getenv("PORT")
	db := mysql.NewMySqlDB("db4free.net","hcs123","hcs12345","3306","restaurants_123")
	defer db.Close()
	s := server.NewServer(db)
	s.Start(":"+port)
}
