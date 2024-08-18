package service

import "github.com/Meduzz/dsl/api"

func NewService(name string, kind ServiceKind) *Service {
	s := &Service{}

	s.Name = name
	s.Kind = kind

	return s
}

func (s *Service) AddVolumes(volume ...string) {
	s.Volumes = append(s.Volumes, volume...)
}

func (s *Service) API() *api.Api {
	if s.Api != nil {
		return s.Api
	}

	a := &api.Api{}
	s.Api = a
	return a
}

func (s *Service) TCP(port int) *Port {
	p := &Port{}

	p.Port = port
	p.Protocol = "tcp"

	s.Ports = append(s.Ports, p)

	return p
}

func (s *Service) UDP(port int) *Port {
	p := &Port{}

	p.Port = port
	p.Protocol = "udp"

	s.Ports = append(s.Ports, p)

	return p
}

func (s *Service) Argv(name string) *Config {
	p := &Config{}

	p.Name = name
	p.Kind = Argument

	s.Params = append(s.Params, p)

	return p
}

func (s *Service) Env(name string) *Config {
	p := &Config{}

	p.Name = name
	p.Kind = Environment

	s.Params = append(s.Params, p)

	return p
}
