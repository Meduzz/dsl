package service

import (
	"github.com/Meduzz/dsl/api"
	"github.com/Meduzz/dsl/deploy"
	"github.com/Meduzz/dsl/proxy"
)

func NewService(name string, kind ServiceKind) *Service {
	s := &Service{}
	s.Name = name
	s.Kind = kind
	return s
}

func (s *Service) API() *api.Api {
	if s.Api != nil {
		return s.Api
	}

	a := &api.Api{}
	s.Api = a
	return a
}

func (s *Service) DeployConfig(image string) *deploy.Deploy {
	s.Deploy = deploy.DeployConfig(image)

	return s.Deploy
}

func (s *Service) ProxyConfig(domain, context string) *proxy.Proxy {
	s.Proxy = proxy.ProxyConfig(domain, context)
	return s.Proxy
}

func (s *Service) Annotations(annotations ...string) *Service {
	s.Tags = append(s.Tags, annotations...)
	return s
}
