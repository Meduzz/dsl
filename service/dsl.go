package service

import (
	"reflect"

	"github.com/Meduzz/dsl/deploy"
)

func NewService(name string, kind ServiceKind) *Service {
	s := &Service{}

	s.Name = name
	s.Kind = kind

	return s
}

func (s *Service) AddEndpoint(endpoint *Endpoint) *Endpoint {
	s.Endpoints = append(s.Endpoints, endpoint)
	return endpoint
}

func (s *Service) AddPort(port *Port) *Port {
	s.Ports = append(s.Ports, port)
	return port
}

func (s *Service) AddParam(param *Config) *Config {
	s.Params = append(s.Params, param)
	return param
}

func (s *Service) SetDeploy(deploy *deploy.Deploy) *deploy.Deploy {
	s.Deploy = deploy

	return deploy
}

func GET(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "GET"

	return e
}

func POST(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "POST"

	return e
}

func PUT(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "PUT"

	return e
}

func DELETE(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "DELETE"

	return e
}

func PATCH(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "PATCH"

	return e
}

func OPTION(path string) *Endpoint {
	e := &Endpoint{}

	e.Path = path
	e.Method = "OPTION"

	return e
}

func (e *Endpoint) AddArgument(param *Param) *Param {
	e.Arguments = append(e.Arguments, param)
	return param
}

func TCP(port int) *Port {
	p := &Port{}

	p.Port = port
	p.Protocol = "tcp"

	return p
}

func UDP(port int) *Port {
	p := &Port{}

	p.Port = port
	p.Protocol = "udp"

	return p
}

func (p *Port) ToMapping(host int) *deploy.PortMap {
	return deploy.NewPortMap(p.Protocol, p.Port, host)
}

func Argv(name string) *Config {
	p := &Config{}

	p.Name = name
	p.Kind = Argument

	return p
}

func Env(name string) *Config {
	p := &Config{}

	p.Name = name
	p.Kind = Environment

	return p
}

func (c *Config) ToConfigData(value string) *deploy.ConfigData {
	return deploy.NewConfigData(c.Name, value, deploy.ConfigKind(c.Kind))
}

func PathVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = PathKind

	return p
}

func QueryVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = QueryKind

	return p
}

func BodyVariable(name, contentType string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = BodyKind
	p.Format = contentType

	return p
}

func HeaderVariable(name string) *Param {
	p := &Param{}

	p.Name = name
	p.Kind = HeaderKind

	return p
}

func (p *Param) SetType(it any) {
	v := reflect.ValueOf(it)

	if v.Kind() == reflect.Pointer {
		println("it's a pointer")
		p.Pointer = true
		v = v.Elem() // drop pointer
	}

	t := v.Type()
	p.Type = t.String()
}

func (p *Param) ArrayOf(it any) {
	p.Array = true
	p.SetType(it)
}

func (p *Param) MapOf(it any) {
	p.Map = true
	p.SetType(it)
}
