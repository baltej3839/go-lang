package main

import "fmt"


type Server struct {
	Host     string
	Port     int
	TLS      bool
	Timeout  int 
	MaxConn  int 
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	s:=&Server{
		Host: "localhost",
        Port: 8080,
        TLS: false,
        Timeout: 10,
	}

	for _, opt:= range opts {
		opt(s)
	}

	return s 
}

func WithTimeOut(timeout int) Option {
	return func(s *Server) {
		s.Timeout=timeout
	}
}


func WithMaxConn(n int) Option {
	return func(s *Server) {
		s.MaxConn=n
	}
}


func WithTLS(enabled bool) Option {
    return func(s *Server) {
        s.TLS = enabled
    }
}



func main() {
	server:=NewServer(
				WithTimeOut(23),
				WithTLS(true),
				WithMaxConn(20),
			)
	fmt.Println(server.Host)
}