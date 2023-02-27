package server

import (
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	portNum int
}

func NewServer(startingPort int) *Server {
	return &Server{
		portNum: startingPort,
	}
}

func (s *Server) Start() error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	attempts := 0
	for ; attempts < 7; attempts++ {
		listenPort := fmt.Sprintf(":%d", s.portNum)
		fmt.Printf("Attempting to start server on port %d\n", s.portNum)
		if err := http.ListenAndServe(listenPort, nil); err != nil {
			fmt.Println(err.Error())
			s.portNum++
		}
	}
	if attempts == 7 {
		return errors.New("Unable to start server, exiting program")
	}
	return nil
}
