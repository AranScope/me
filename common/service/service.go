package service

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

		go (func() {
			http.Handle("/metrics", promhttp.Handler())
			err := http.ListenAndServe(":2112", nil)
			if err != nil {
				panic(err)
			}
		})()

		err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router)
		if err != nil {
			panic(err)
		}
	}
}
