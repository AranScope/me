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

var svc *Service

func Name() string {
	if svc != nil {
		return svc.Name
	}
	return "uninitialised"
}

func Init(svcName string) *Service {
	svc = &Service{
		Name: svcName,
	}

	return svc
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
