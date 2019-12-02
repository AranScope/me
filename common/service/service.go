package service

import (
	"fmt"
	"net/http"
)

type Service struct {
	Name   string
	Router http.Handler
	Port   int
}

func Init(svcName string) *Service {
	// todo: configure prometheus metrics here or something? Or do we want to do envoy here? What does monzo do in bedrock?
	// todo: how can I make this more transport independent?
	//  - generate request clients based on some schema definitions for endpoints?
	// 		- protobuf?
	//		- yaml?
	svc := Service{
		Name: svcName,
	}
	return &svc
}

func (s *Service) WithRouter(port int, router http.Handler) *Service {
	s.Port = port
	s.Router = router
	return s
}

func (s *Service) Start() {
	if s.Router != nil {
		err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router)
		if err != nil {
			panic(err)
		}
	}
}
