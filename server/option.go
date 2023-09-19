package server

import (
	"net"
	"strings"
	"time"
)

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		host := strings.Split(s.server.Addr, ":")[0]
		s.server.Addr = net.JoinHostPort(host, port)
	}
}

func Host(host string) Option {
	return func(s *Server) {
		port := strings.Split(s.server.Addr, ":")[1]
		s.server.Addr = net.JoinHostPort(host, port)
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = time.Second * timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = time.Second * timeout
	}
}
