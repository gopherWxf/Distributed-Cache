package tcp

import (
	"Distributed-cache/cache"
	"log"
	"net"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", ":12346")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go s.process(conn)
	}
}
func New(cache cache.Cache) *Server {
	return &Server{cache}
}
