package server

import (
	"github.com/restaurant/Res_Man_MicroService/pkg/database"
)

type Server struct {
	db     database.Database
	router *Router
}

func NewServer(db database.Database) *Server {
	return &Server{
		db:     db,
		router: NewRouter(db),
	}
}

func (s *Server) Start(port string) {
	engine := s.router.Router()
	if err := engine.Run(port); err != nil {
		panic(err)
	}
}
